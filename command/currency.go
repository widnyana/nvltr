package command

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/leekchan/accounting"
	"github.com/redite/tlbot"
)

func init() {
	register(cmdCurrency)
}

var cmdCurrency = &Command{
	Name:      "cur",
	ShortLine: "currency converter [/cur FROM TO AMOUNT | /cur USD IDR 666]",
	Run:       runCurrency,
}

var (
	baseURL = "http://www.currency-converter.org.uk/widget/CCUK-CC2-AJAX.php?ConvertTo=%s&ConvertFrom=%s&amount=%s"
)

func runCurrency(ctx context.Context, b *tlbot.Bot, msg *tlbot.Message) {
	opts := new(tlbot.SendOptions)
	opts.ReplyTo = msg.ID

	args := msg.Args()

	if len(args) < 3 || len(args) > 3 {
		b.SendMessage(msg.Chat.ID, "format: `/cur FROM TO AMOUNT`.", opts)
		return
	}

	from := strings.ToUpper(args[0])
	to := strings.ToUpper(args[1])
	amount := args[2]
	floatAmount, _ := strconv.ParseFloat(amount, 64)
	acfrom := accounting.FormatNumber(floatAmount, 2, ".", ",")

	url := fmt.Sprintf(
		baseURL,
		to,
		from,
		amount,
	)
	res, err := httpclient.Get(url)
	if err != nil {
		b.SendMessage(msg.Chat.ID, err.Error(), opts)
		return
	}
	defer res.Body.Close()

	result, _ := ioutil.ReadAll(res.Body)
	newresm, _ := strconv.ParseFloat(string(result), 64)
	ac := accounting.Accounting{Symbol: to + " ", Precision: 2, Decimal: ",", Thousand: "."}
	money := ac.FormatMoneyBigFloat(big.NewFloat(newresm))

	message := fmt.Sprintf("%s %s to %s => %s", acfrom, from, to, money)

	_, err = b.SendMessage(msg.Chat.ID, message, opts)
	if err != nil {
		log.Printf("Error while sending message. Err: %v\n", err)
	}
}
