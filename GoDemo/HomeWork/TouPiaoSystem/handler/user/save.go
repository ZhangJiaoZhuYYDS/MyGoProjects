package user

//func Save(formInfo *model.VotesVotesopt , value any) error {
//	var votesOpt model.Votesopt
//	s := ""
//	for i := range formInfo.Hobbies {
//		s+= " "+formInfo.Hobbies[i]
//	}
//	fmt.Println(s)
//	votesOpt.Hobbies = s
//	votesOpt.Id = cast.ToUint64(value)
//	votesOpt.Sex = formInfo.Sex
//	vates := model.Votes{
//		BaseModel:   model.BaseModel{
//		},
//		Title:       "测试",
//		VotesOpt:    []model.VotesOpt{votesOpt},
//		TotalNumber: 0,
//	}
//	err := model.DB.Self.Debug().Create(&vates).Error
//	if err != nil {
//		log.Println("创建投票表失败")
//		return err
//	}
//	log.Println("sucesss")
//	return nil
//}
