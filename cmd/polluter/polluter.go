package main

import (
	"context"
	"flag"
	"log"

	"miniWiki"
	cQuery "miniWiki/domain/category/query"
	rQuery "miniWiki/domain/resource/query"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	cfg, err := miniWiki.InitConfig()
	if err != nil {
		panic(err)
	}
	pool, _ := pgxpool.Connect(context.Background(), cfg.Database.DatabaseUrl)
	ctx := context.Background()
	categoryQuerier := cQuery.NewQuerier(pool)
	resourceQuerier := rQuery.NewQuerier(pool)

	var n int

	flag.IntVar(&n, "n", 50, "number of data entries generated")

	for i := 0; i < n/2; i++ {
		_, err := categoryQuerier.InsertCategory(ctx, gofakeit.BuzzWord())
		log.Println(err)
	}

	for i := 0; i < n/2; i++ {
		_, err := categoryQuerier.InsertSubCategory(ctx, gofakeit.BuzzWord(), gofakeit.Number(1, n/2))
		log.Println(err)
	}

	for i := 0; i < n/2; i++ {
		_, err := resourceQuerier.InsertResource(ctx, rQuery.InsertResourceParams{
			Title:       gofakeit.LoremIpsumWord(),
			Description: gofakeit.LoremIpsumSentence(15),
			Link:        gofakeit.URL(),
			CategoryID:  gofakeit.Number(1, n),
		})
		log.Println(err)
	}

	pool.Close()
}
