package main

import (
	"context"
	"flag"
	"log"

	"miniWiki/config"
	cQuery "miniWiki/domain/category/query"
	rQuery "miniWiki/domain/resource/query"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	pool, err := pgxpool.Connect(context.Background(), cfg.Database.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	ctx := context.Background()
	categoryQuerier := cQuery.NewQuerier(pool)
	resourceQuerier := rQuery.NewQuerier(pool)

	var n int

	flag.IntVar(&n, "n", 50, "number of data entries generated")

	for i := 0; i < n/2; i++ {
		_, err := categoryQuerier.InsertCategory(ctx, gofakeit.BuzzWord())
		if err != nil {
			log.Println(err)
		}
	}

	for i := 0; i < n/2; i++ {
		_, err := categoryQuerier.InsertSubCategory(ctx, gofakeit.LoremIpsumSentence(2), gofakeit.Number(1, n/2))
		if err != nil {
			log.Println(err)
		}
	}

	for i := 0; i < n/2; i++ {
		_, err := resourceQuerier.InsertResource(ctx, rQuery.InsertResourceParams{
			Title:       gofakeit.LoremIpsumSentence(2),
			Description: gofakeit.LoremIpsumSentence(15),
			Link:        gofakeit.URL(),
			CategoryID:  gofakeit.Number(1, n),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
