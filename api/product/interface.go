package product

type PageNum struct {
	Name    string `json:"name"`
	PageNum string `json:"page_num"`
}

type Sku struct {
	Id                    uint64 `json:"sku,omitempty"`          // 商品编号
	Name                  string `json:"name,omitempty"`         // 商品名称
	Introduction          string `json:"introduction,omitempty"` // 商品详情页大字段
	SaleUnit              string `json:"saleUnit,omitempty"`     // 售卖单位
	Weight                string `json:"weight,omitempty"`       // 重量
	ProductArea           string `json:"productArea,omitempty"`  // 产地
	WareQD                string `json:"wareQD,omitempty"`       // 包装清单
	ImagePath             string `json:"imagePath,omitempty"`    // 主图
	Param                 string `json:"param,omitempty"`        // 规格参数
	State                 int    `json:"state,omitempty"`        // 状态
	BrandName             string `json:"brandName,omitempty"`    // 品牌名称
	Upc                   string `json:"upc,omitempty"`          // UPC码, 区分实物、图书、音像、三种场景
	Category              string `json:"category,omitempty"`     // 分类, 示例"670;729;4837"
	Brand                 string `json:"Brand,omitempty"`        // 品牌
	Sheet                 string `json:"Sheet,omitempty"`        // 印张
	RelatedProducts       string `json:"relatedProducts,omitempty"`
	ISBN                  string `json:"ISBN,omitempty"`    // ISBN
	Editer                string `json:"Editer,omitempty"`  // 编者
	PrintNo               string `json:"PrintNo,omitempty"` // 印次
	Author                string `json:"Author,omitempty"`  // 作者
	PackNum               string `json:"PackNum,omitempty"` // 套装数量
	ContentDesc           string `json:"contentDesc,omitempty"`
	PrintTime             string `json:"PrintTime,omitempty"`   // 印刷时间
	SkuType               string `json:"skuType,omitempty"`     // 类型（例book）
	Papers                string `json:"Papers,omitempty"`      // 用纸
	Package               string `json:"Package,omitempty"`     // 包装(装帧)
	Proofreader           string `json:"Proofreader,omitempty"` // 校对
	EditerDesc            string `json:"editerDesc,omitempty"`
	BookAbstract          string `json:"bookAbstract,omitempty"`
	Catalogue             string `json:"catalogue,omitempty"`
	PublishTime           string `json:"PublishTime,omitempty"` // 出版时间
	Pages                 string `json:"Pages,omitempty"`       // 页数
	AuthorDesc            string `json:"AuthorDesc,omitempty"`
	Image                 string `json:"image,omitempty"`    // 图片
	Transfer              string `json:"Transfer,omitempty"` // 译者
	Appintroduce          string `json:"appintroduce,omitempty"`
	Drawer                string `json:"Drawer,omitempty"`    // 绘者
	Language              string `json:"Launguage,omitempty"` // 图书语言
	BatchNo               string `json:"BatchNo,omitempty"`   // 版次
	Comments              string `json:"Comments,omitempty"`
	Press                 string `json:"Press,omitempty"`       // 出版社
	Foreignname           string `json:"Foreignname,omitempty"` // 外文名
	Format                string `json:"Format,omitempty"`      // 格式
	Performer             string `json:"Performer,omitempty"`   // 演奏者
	Soundtrack            string `json:"Soundtrack,omitempty"`  // 碟数
	Actor                 string `json:"Actor,omitempty"`       // 演员
	Dregion               string `json:"Dregion,omitempty"`     // 地区
	Voiceover             string `json:"Voiceover,omitempty"`   // 解说者
	Director              string `json:"Director,omitempty"`    // 导演
	BoxContents           string `json:"box_Contents,omitempty"`
	LanguageSubtitled     string `json:"Language_Subtitled,omitempty"`     // 字幕语言
	Media                 string `json:"Media,omitempty"`                  // 介质
	ScreenRatio           string `json:"Screen_Ratio,omitempty"`           // 屏幕比例
	Episode               string `json:"Episode,omitempty"`                // 集数
	MvdWxjz               string `json:"Mvd_Wxjz,omitempty"`               // 文像进字
	PublishingCompany     string `json:"Publishing_Company,omitempty"`     // 发行公司
	ISRC                  string `json:"ISRC,omitempty"`                   // ISRC
	Singer                string `json:"Singer,omitempty"`                 // 演唱者
	LanguagePronunciation string `json:"Language_Pronunciation,omitempty"` // 发音语言
	ProductionCompany     string `json:"Production_Company,omitempty"`     // 出品公司
	AudioEncodingChinese  string `json:"Audio_Encoding_Chinese,omitempty"` // 音频格式
	Authors               string `json:"Authors,omitempty"`                // 作词
	Aka                   string `json:"Aka,omitempty"`                    // 又名
	Region                string `json:"Region,omitempty"`                 // 区码
	Copyright             string `json:"Copyright,omitempty"`              // 版权提供
	Compose               string `json:"Compose,omitempty"`                // 作曲
	Screenwriter          string `json:"Screenwriter,omitempty"`           // 编剧
	LanguageDubbed        string `json:"Language_Dubbed,omitempty"`        // 配音语言
	Manual                string `json:"manual,omitempty"`
	Length                string `json:"Length,omitempty"` // 片长
	MaterialDescription   string `json:"material_Description,omitempty"`
	ReleaseDate           string `json:"ReleaseDate,omitempty"`
	WxIntroduction        string `json:"wxintroduction,omitempty"`
}

type SkuImage struct {
	Id        uint64 `json:"id,omitempty"`
	SkuId     uint64 `json:"skuId,omitempty"`
	Path      string `json:"path,omitempty"` // 前缀：http://img13.360buyimg.com/n0/, 其中 n0(最大图 800*800px)、n1(350*350px)、n2(160*160px)、n3(130*130px)、n4(100*100px)为图片大小。注意：n0带有京东水印，其余的n1-n4不带，12 大图无水印的 。
	Created   string `json:"created,omitempty"`
	Modified  string `json:"modified,omitempty"`
	Yn        uint   `json:"yn,omitempty"`        // 0:不可用;1:可用
	IsPrimary uint   `json:"isPrimary,omitempty"` // 是否主图 1：是 0：否
	OrderSort uint   `json:"orderSort,omitempty"` // 排序
	Position  int    `json:"position,omitempty"`  // 位置
	Type      uint   `json:"type,omitempty"`      // 类型（0方图  1长图【服装】）
	Features  string `json:"features,omitempty"`  // 特征

}

type SkuState struct {
	State uint   `json:"state,omitempty"` // 1：上架，0：下架
	SkuId uint64 `json:"sku,omitempty"`
}

type SkuSaleState struct {
	SkuId            uint64 `json:"skuId,omitempty"`
	Name             string `json:"name,omitempty"`
	SaleState        uint   `json:"saleState,omitempty"`        // 是否可售，1：是，0：否
	IsCanVAT         uint   `json:"isCanVAT,omitempty"`         // 是否可开专票，1：支持，0：不支持
	NoReasonToReturn *uint  `json:"noReasonToReturn,omitempty"` // 无理由退货类型：0,1,2,3,4,5,6,7,8; 0、3：不支持7天无理由退货； 1、5、8或null：支持7天无理由退货；2：支持90天无理由退货；4、7：支持15天无理由退货；6：支持30天无理由退货；（提示客户取到其他枚举值，无效）
	Thwa             uint   `json:"thwa,omitempty"`             // 无理由退货文案类型：null：文案空；0：文案空；1：支持7天无理由退货；2：支持7天无理由退货（拆封后不支持）；3：支持7天无理由退货（激活后不支持）；4：支持7天无理由退货（使用后不支持）；5：支持7天无理由退货（安装后不支持）；12：支持15天无理由退货；13：支持15天无理由退货（拆封后不支持）；14：支持15天无理由退货（激活后不支持）；15：支持15天无理由退货（使用后不支持）；16：支持15天无理由退货（安装后不支持）；22：支持30天无理由退货；23：支持30天无理由退货（安装后不支持）；24：支持30天无理由退货（拆封后不支持）；25：支持30天无理由退货（使用后不支持）；26：支持30天无理由退货（激活后不支持）；（提示客户取到其他枚举值，无效）
	IsSelf           uint   `json:"isSelf,omitempty"`
	Is7ToReturn      uint   `json:"is7ToReturn,omitempty"`
}

type SkuAreaLimit struct {
	SkuId          uint64 `json:"skuId,omitempty"`
	IsAreaRestrict bool   `json:"isAreaRestrict,omitempty"`
}

type SkuGifts struct {
	Gifts          []SkuGift `json:"gifts,omitempty"`          // 赠品附件列表
	MaxNum         uint      `json:"maxNum,omitempty"`         // 赠品要求最多购买数量（为0表示没配置）
	MinNum         uint      `json:"minNum,omitempty"`         // 赠品要求最少购买数量 （为0表示没配置）
	PromoStartTime int64     `json:"promoStartTime,omitempty"` // 促销开始时间
	PromoEndTime   int64     `json:"promoEndTime,omitempty"`   // 促销结束时间
}

type SkuGift struct {
	SkuId uint64 `json:"skuId,omitempty"`
	Num   uint   `json:"num,omitempty"`
	Type  uint   `json:"giftType,omitempty"` // 1：附件，2：赠品。附件是指与主商品配套使用的部分，如空调的外机。赠品是不影响主商品使用的附赠商品。下单时，可以选择是否要赠品，但附件默认都必须要。
}

type Yanbao struct {
	MainSkuId    uint64         `json:"mainSkuId,omitempty"`         // 主商品的sku
	ImgUrl       string         `json:"imgUrl,omitempty"`            // 保障服务类别显示图标url
	DetailUrl    string         `json:"detailUrl,omitempty"`         // 保障服务类别静态页详情url
	DisplayNo    int            `json:"displayNo,omitempty"`         // 保障服务类别显示排序
	CategoryCode string         `json:"categoryCode,omitempty"`      // 保障服务分类编码
	DisplayName  string         `json:"displayName,omitempty"`       // 保障服务类别名称
	List         []YanbaoDetail `json:"fuwuSkuDetailList,omitempty"` // 保障服务商品详情列表
}

type YanbaoDetail struct {
	BindSkuId   uint64  `json:"bindSkuId,omitempty"`   // 保障服务skuId
	BindSkuName string  `json:"bindSkuName,omitempty"` // 保障服务sku名称（6字内）
	SortIndex   int     `json:"sortIndex,omitempty"`   // 显示排序
	Price       float64 `json:"price,omitempty"`       // 保障服务sku价格
	Tip         string  `json:"tip,omitempty"`         // 保障服务说明提示语（20字内）
	Favor       bool    `json:"favor,omitempty"`       // 是否是优惠保障服务（PC单品页、PC购物车会根据此标识是否展示优惠图标，优惠图标单品页提供）
}

type SkuIsCod struct {
	PayFirst bool   `json:"payFirst,omitempty"`  // 是否支持货到付款
	Reason   string `json:"nonCODmsg,omitempty"` // 错误原因
}

type SimilarProduct struct {
	Dim          int        `json:"dim,omitempty"`          // 维度
	SaleName     string     `json:"saleName,omitempty"`     // 销售名称
	SaleAttrList []SaleAttr `json:"saleAttrList,omitempty"` // 商品销售标签;销售属性下可能只有一个标签，此种场景可以选择显示销售名称和标签，也可以不显示
}

type SaleAttr struct {
	ImagePath string   `json:"imagePath,omitempty"` // 标签图片地址
	SaleValue string   `json:"saleValue,omitempty"` // 标签名称
	SkuIds    []uint64 `json:"skuIds,omitempty"`    // 当前标签下的同类商品skuId
}

type Category struct {
	Id       uint64 `json:"catId,omitempty"`
	ParentId uint64 `json:"parentId,omitempty"`
	Name     string `json:"name,omitempty"`
	Class    uint   `json:"catClass,omitempty"` // 0：一级分类；1：二级分类；2：三级分类；
	State    uint   `json:"state,omitempty"`    // 1：有效；0：无效；
}

type CategoriesResult struct {
	Total      int        `json:"totalRows,omitempty"`
	PageNo     int        `json:"pageNo,omitempty"`
	PageSize   int        `json:"pageSize,omitempty"`
	Categories []Category `json:"categorys,omitempty"`
}
