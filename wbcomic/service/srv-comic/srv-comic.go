package srv_comic

import(
	"context"

	"wbcomic/pb"
	"wbcomic/core"
	"time"
)

type ComicInfo struct {
	cc 					pb.ComicServiceClient 	`xorm:"-"`
	ComicId             int64    `xorm:"not null pk autoincr INT(11)"`
	ComicName           string   `xorm:"not null VARCHAR(50)"`
	Status              int32    `xorm:"not null default 1 TINYINT(3)"`
	CreateTime          int64    `xorm:"not null index INT(10)"`
}


func (c *ComicInfo) GetComicRowById (ctx context.Context, in *pb.ComicFilterRequest) (*pb.ComicRowResponse, error){

	comicRow := new(ComicInfo)
	_, err := core.MasterDB("comic").Id(in.ComicId).Get(comicRow)

	return &pb.ComicRowResponse{
		ComicId:comicRow.ComicId,
		ComicName:comicRow.ComicName,
		Status:comicRow.Status,
		CreateTime:comicRow.CreateTime,
	},err

}

func (c *ComicInfo) GetComicListByPage (ctx context.Context, in *pb.ComicFilterRequest) (comicListResponse *pb.ComicListResponse, err error){

	limit := int(in.RowNum)
	start := int((in.PageNum - 1) * in.RowNum)

	comicList := make([]ComicInfo, 0)
	err = core.MasterDB("comic").Where("status = ?", 1).Limit(limit, start).Find(&comicList)

	time.Sleep(3*time.Second)
//panic("aaaaa3333")
	comicListResponse = new(pb.ComicListResponse)

	for _,v := range comicList{
		comicListResponse.Data = append(comicListResponse.Data, &pb.ComicRowResponse{
			v.ComicId,
			v.ComicName,
			v.CreateTime,
			v.Status,
		})
	}

	return comicListResponse,err
}

func (c *ComicInfo)AddComic(ctx context.Context, in *pb.ComicRowResponse) (*pb.AddComicResponse, error)  {

	comicRow := new(ComicInfo)
	comicRow.ComicName = in.ComicName
	comicRow.Status = in.Status
	comicRow.CreateTime = in.CreateTime

	_, err := core.MasterDB("comic").Insert(comicRow)

	return &pb.AddComicResponse{
		ComicId:comicRow.ComicId,
	},err
}