package handler

import (
	"app/cmd/app/infrastructure"
	"app/cmd/app/service"
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func LineHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		os.Getenv("LINE_SECRET"),
		os.Getenv("LINE_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		user := infrastructure.FindByLineId(event.Source.UserID)
		if event.Type == linebot.EventTypePostback {

			data := event.Postback.Data
			values, err := url.ParseQuery(data)
			if err != nil {
				log.Println(err)
			}

			distance, err := strconv.ParseInt(values.Get("distance"), 10, 32)
			if err != nil {
				log.Println((err))
			}
			user.Distance = sql.NullInt32{Int32: int32(distance), Valid: true}
		}
		if event.Type == linebot.EventTypeMessage {
			if err != nil {
				log.Println(err)
			}

			switch message := event.Message.(type) {

			case *linebot.TextMessage:
				log.Print(message)

			case *linebot.StickerMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1", "13")).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				user.Latitude = sql.NullFloat64{Float64: message.Latitude, Valid: true}
				user.Longitude = sql.NullFloat64{Float64: message.Longitude, Valid: true}
			}
		}
		user.LineID = event.Source.UserID
		infrastructure.Save(user)
		if !user.Distance.Valid {
			data, err := ioutil.ReadFile("templates/flex_dialog.json")

			if err != nil {
				log.Println(err)
			}

			container, err := linebot.UnmarshalFlexMessageJSON(data)
			if err != nil {
				log.Println(err)
			}

			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("distance please", container)).Do(); err != nil {
				log.Print(err)
			}
		} else if !user.Longitude.Valid || !user.Latitude.Valid {
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("location please")).Do(); err != nil {
				log.Print(err)
			}
		} else {
			destination := service.CalcDistance(user.Latitude.Float64, user.Longitude.Float64, user.Distance.Int32)
			log.Println(destination)

			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewLocationMessage("Today's Destication", "unknown", destination.Latitude, destination.Longitude)).Do(); err != nil {
				log.Print(err)
			}
		}
	}
}
