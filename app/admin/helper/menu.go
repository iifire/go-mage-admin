package helper

import (
	"fmt"
	"github.com/sbabiv/xml2map"
	"go-mage-admin/app/mage"
	"log"
	"strings"
)
import "github.com/beevik/etree"

type Menu struct {
}

// GetMenus 获取admin菜单
func (menu *Menu) GetMenus() map[string]interface{} {
	//TODO... save cache
	var xmlStr string
	var docAppendArr = make([]*etree.Element, 0)
	//log.Println("AppModules=", mage.AppModules)
	docNew := etree.NewDocument()
	docNewMenu := docNew.CreateElement("menu")
	for k, _ := range mage.AppModules {
		path := "./app/" + k + "/etc/admin.xml"
		if mage.Exists(path) {
			doc := etree.NewDocument()
			if err := doc.ReadFromFile(path); err != nil {
				panic(err)
			}
			for _, item := range doc.SelectElement("menu").ChildElements() {
				log.Println("item.Tag=", item.Tag)
				if item.Tag == "append" {
					for _, item2 := range item.ChildElements() {
						docAppendArr = append(docAppendArr, item2)
					}
				} else {
					docNewMenu.AddChild(item)
				}
			}
		}
	}
	//log.Println("docAppendArr=", docAppendArr)
	// 追加菜单处理
	for _, item := range docAppendArr {
		if item.SelectElement("to") == nil {
			continue
		}
		toArr := strings.Split(item.SelectElement("to").Text(), "/")
		//父级菜单最多三层
		toLen := len(toArr)
		if toLen > 0 {
			p1 := docNewMenu.SelectElement(toArr[0])
			if p1 != nil {
				p1Child := p1.SelectElement("children")
				if p1Child == nil {
					p1Child = p1.CreateElement("children")
				}
				if toLen == 1 {
					p1Child.AddChild(item)
				} else {
					p2 := p1Child.SelectElement(toArr[1])
					if p2 != nil {
						p2Child := p2.SelectElement("children")
						if p2Child == nil {
							p2Child = p2.CreateElement("children")
						}
						if toLen == 2 {
							p2Child.AddChild(item)
						} else {
							p3 := p2Child.SelectElement(toArr[2])
							if p3 != nil {
								p3Child := p3.SelectElement("children")
								if p3Child == nil {
									p3Child = p3.CreateElement("children")
								}
								if toLen == 3 {
									p3Child.AddChild(item)
								}
							}
						}
					}
				}
			}
		}
	}
	xmlStr, err := docNew.WriteToString()
	//log.Println("xml=", xmlStr)
	decoder := xml2map.NewDecoder(strings.NewReader(xmlStr))
	menus, err := decoder.Decode()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	//log.Println("menus", menus)
	return menus
}
