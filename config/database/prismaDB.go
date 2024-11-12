package config

import (
	"go-websocket-fiber/prisma/db"
	"log"
)

var prismaSQL *db.PrismaClient  


func InitPrismaDB() error {
	prismaSQL = db.NewClient()
	if err := prismaSQL.Prisma.Connect(); err != nil {
		log.Println("PrismaDB Database connection fell")
		return err
	} else{
		log.Println("PrismaDB Database connection established")
	}
	return nil
}

// GetPrismaDB returns the Prisma client instance
func GetPrismaDBClient() *db.PrismaClient {
	return prismaSQL
}