package main

import (
	"log"

	"github.com/bububa/jdvop"
	// "github.com/bububa/jdvop/api/area"
	"github.com/bububa/jdvop/api/product"
	//"github.com/bububa/jdvop/api/search"
	//"github.com/bububa/jdvop/api/price"
	//"github.com/bububa/jdvop/api/stock"
)

func main() {
	client := jdvop.NewClient("xxx", "xxx", "xxx", "xxx")
	ret, err := product.GetDetail(client, "54659086878", "nappintroduction,wxintroduction")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", ret)
	/*
	   skuNums := []stock.SkuNum{
	       {
	           SkuId: 202551,
	           Num:   10,
	       },
	   }
	   ret, err := stock.GetNewStockById(client, skuNums, "13_1032_1033_0")
	   if err != nil {
	       log.Fatalln(err)
	   }
	   log.Printf("%+v\n", ret)
	*/
	/*
		ret, err := price.GetSellPrice(client, "266910", "Price,marketPrice,containsTax,nakedPrice")
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%+v\n", ret)
	*/
	/*
		req := &search.SearchRequest{
			Keyword: "macbook",
		}
		ret, err := req.Search(client)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%+v\n", ret)
	*/
	/*
		        ret, err := product.GetCategory(client, 670)
		        if err != nil {
		            log.Fatalln(err)
		        }
		        log.Printf("%+v\n", ret)
		        ret, err := product.GetSimilarSku(client, "266910")
		        if err != nil {
		            log.Fatalln(err)
		        }
		        log.Printf("%+v\n", ret)
			    ret, err := product.Check(client, "266910", "noReasonToReturn,thwa,isSelf,isJDLogistics,taxInfo")
			    if err != nil {
			        log.Fatalln(err)
			    }
			    log.Printf("%+v\n", ret)
			    ret, err := product.GetSkuState(client, "266910")
			    if err != nil {
			        log.Fatalln(err)
			    }
			    log.Printf("%+v\n", ret)
			    ret, err := product.GetSkuImage(client, "266910")
			    if err != nil {
			        log.Fatalln(err)
			    }
			    log.Printf("%+v\n", ret)
			    ret, err := product.GetDetail(client, "266910", "nappintroduction")
			    if err != nil {
			        log.Fatalln(err)
			    }
			    log.Printf("%+v\n", ret)
			    ret, err := product.GetSkuByPage(client, "16", 1)
			    if err != nil {
			        log.Fatalln(err)
			    }
			    log.Printf("%+v\n", ret)
		        pages, err := product.GetPageNum(client, "")
		        if err != nil {
		            log.Fatalln(err)
		        }
			    log.Printf("%+v\n", pages)
	*/
	/*
	       addr, err := area.GetJDAddressFromAddress(client, "北京市大兴区亦庄经济开发区经海路5号")
	       if err != nil {
	           log.Fatalln(err)
	       }
	       log.Printf("%+v\n", addr)
	   	provinces, err := area.GetProvince(client)
	   	if err != nil {
	   		log.Fatalln(err)
	   	}
	   	log.Printf("%+v\n", provinces)
	   	cities, err := area.GetCity(client, 19)
	   	if err != nil {
	   		log.Fatalln(err)
	   	}
	   	log.Printf("%+v\n", cities)
	   	counties, err := area.GetCounty(client, 1000)
	   	if err != nil {
	   		log.Fatalln(err)
	   	}
	   	log.Printf("%+v\n", counties)
	   	towns, err := area.GetTown(client, 46666)
	   	if err != nil {
	   		log.Fatalln(err)
	   	}
	   	log.Printf("%+v\n", towns)
	*/
}
