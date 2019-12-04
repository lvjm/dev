package main

//导入包
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
	"strings"
)

//全局变量
const (
	CourtFileInfoCertUnArchived = "0"
	CourtFileInfoCertArchived   = "1"
	compositekeyIdIndexName     = "originalFileHash~originalTxHash~fileHash~txHash"
	compositeIndexName          = "Event~bizId~eventName~operator~timestamp"
	//invoke
	NotFoundFuncErrStr = "没有找到此方法："
	UndefinedFunStr    = "调用了未定义函数！"
	//通用
	TxTimestampErrStr                   = "获取当前时间失败！"
	PutStateErrStr                      = "上链失败！"
	GetStateErrStr                      = "从链上获取数据失败！"
	GetStateByPartialCompositeKeyErrStr = "根据CompositeKey获取数据失败！"
	CreateCompositeKeyErrStr            = "创建CompositeKey失败！"
	//AddRecord
	AddRecordParameterErrStr       = "参数不正确，期望4个参数，您输入的参数个数为"
	StringToInt64ErrStr            = "creationDate转int64失败!"
	InvalidCreationDateStr         = "CreationDate只能为正数。"
	UnmarshalMetadataErrStr        = "将字符串解析为Metadata异常，"
	MarshalCourtFileInfoCertErrStr = "解析卷宗为byte异常，"
	CreateCourtFileInfoCertErrStr  = "根据输入信息创建卷宗异常！"
	PreFileKeyFormatErrStr         = "PreFileKey格式错误，必须包含'-'字符。"
	FileStatusErrStr               = "添加记录文件状态只能是未归档，状态码为0"
	//GetRecord
	GetRecordParameterErrStr     = "参数不正确，期望2个参数，您输入的参数个数为"
	courtFileInfoCertNotFoundStr = "该文件不存在！"
	//AddEvent
	AddEventParameterErrStr   = "参数不正确，期望4个参数，您输入的参数个数为"
	timeStampStrToInt64ErrStr = "timeStampStr转int64失败!"
	InvalidTimeStampStr       = "timeStampStr只能为正数。"
	//SearchEvent
	SearchEventParameterErrStr = "参数不正确，期望3个参数，您输入的参数个数为"
	EventResultsNextErrStr     = "对eventResults进行遍历异常!"
	MarshalsearchResErrStr     = "将searchRes转为byte失败!"
	//Archive
	ArchiveParameterErrStr = "参数不正确，期望2个参数，您输入的参数个数为"
	//Search
	SearchParameterErrStr    = "参数不正确，期望3个参数，您输入的参数个数为"
	PageSizeParaFormatErrStr = "pageSize参数格式不正确，必须是小于100的正整数！"
	QueryErrStr              = "根据queryString获取数据失败！"
	QueryResultsErrStr       = "遍历queryResults异常！"
	//OriginalFileKeyIdSearch
	OriginalFileKeyIdSearchParameterErrStr = "参数不正确，期望1个参数，您输入的参数个数为"
	GetOriginalFileErrStr                  = "根据KeyId获取原始文件失败！"
	OriginalFileKeyIdSplitLen              = "originalFileKeyId格式不符合要求"
	SplitCompositeKeyErrStr                = "CompositeKey分割异常！"
	InvalidCompositeKey                    = "无效的CompositeKey！"
	originalFileResultsNextErrStr          = "对originalFileResults进行遍历异常！"
	//GetAttestation
	GetAttestationParameterErrStr = "参数不正确，期望3个参数，您输入的参数个数为"
)

//文件事件数据结构
type FileEvent struct {
	Event     string `json:"Event"`     //事件标志
	BizId     string `json:"bizId"`     //主键BizId
	EventName string `json:"eventName"` //事件名称
	Operator  string `json:"operator"`  //操作者
	TimeStamp string `json:"timeStamp"` //时间戳
}

//智能合约
type CourtFileCertChaincode struct {
}
//文件数据结构
type CourtFileInfoCert struct {
	BizId        string   `json:"bizId"`        //文件的key
	ExternalId   string   `json:"externalId"`   //文件bizId
	FileHash     string   `json:"fileHash"`     //文件哈希
	CreationDate int64    `json:"creationDate"` //文件创建日期
	Metadata     Metadata `json:"metadata"`     //文件元数据
}

//分页response数据结构
type ResponseMetadata struct {
	RecordsCount int32  `json:"recordsCount"` //每次查几条数据
	Bookmark     string `json:"bookMark"`     //标记从哪里开始查询
}

//fileEvent查询response数据结构
type SearchEventRes struct {
	FileEvents       []FileEvent      `json:"fileEvents"`       //查询到的文件数组
	ResponseMetadata ResponseMetadata `json:"responseMetadata"` //分页查询返回值
}

//queryString查询response数据结构
type SearchRes struct {
	CourtFileInfoCerts []CourtFileInfoCert `json:"courtFileInfoCerts"` //查询到的文件数组
	ResponseMetadata   ResponseMetadata    `json:"responseMetadata"`   //分页查询返回值
}

//文件元数据数据结构
type Metadata struct {
	FileName       string `json:"fileName"`       //文件名
	Storage        string `json:"storage"`        //文件存储
	DataUri        string `json:"dataUri"`        //文件在分布式系统上的url
	ParentBizId    string `json:"parentBizId"`    //修改前文件哈希
	FileType       string `json:"fileType"`       //文件类型
	Org            string `json:"org"`            //组织机构
	Uploader       string `json:"uploader"`       //操作文件的人
	PersonInCharge string `json:"personInCharge"` //负责人
	Status         string `json:"status"`         //标志文件是否归档
	Description    string `json:"description"`    //文件描述
	MetadataHash   string `json:"metadataHash"`   //元数据哈希
	MimeType       string `json:"mimeType"`       //文件mimeType
}

func main() {
	//主函数中启动智能合约
	err := shim.Start(new(CourtFileCertChaincode))
	if err != nil {
		fmt.Printf("Error starting CourtFileCertChaincode: %s", err)
	}
}

// 初始化方法
func (t *CourtFileCertChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Print("chaincode initial...")
	return shim.Success(nil)
}

// 方法调用中心代理，多个实际的方法调用转发
func (t *CourtFileCertChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//获取合约方法名称和参数
	function, args := stub.GetFunctionAndParameters()
	//判断哪个方法被调用
	if function == "AddRecord" { //调用AddRecord方法
		return t.AddRecord(stub, args)
	} else if function == "GetRecord" { //调用GetRecord方法
		return t.GetRecord(stub, args)
	} else if function == "GetAttestation" { //调用GetAttestation方法
		return t.GetAttestation(stub, args)
	} else if function == "AddEvent" { //调用AddEvent方法
		return t.AddEvent(stub, args)
	} else if function == "SearchEvent" { //调用SearchEvent方法
		return t.SearchEvent(stub, args)
	} else if function == "Archive" { //调用Archive方法
		return t.Archive(stub, args)
	} else if function == "OriginalFileKeyIdSearch" { //调用OriginalFileKeyIdSearch方法
		return t.OriginalFileKeyIdSearch(stub, args)
	} else if function == "Search" { //调用Search方法
		return t.Search(stub, args)
	} else if function == "AttestationMetaDataHash" { //调用AttestationMetaDataHash方法
		return t.AttestationMetaDataHash(stub, args)
	}
	//调用错误处理
	fmt.Printf(NotFoundFuncErrStr + function)
	return shim.Error(UndefinedFunStr)
}

// ===== AddRecord  ========================================================
//  向couchdb中添加一条存证信息
//	@param	externalId fileHash	creationDateStr	metadataJson
//	@return bizId
//	根据入参构造CourtFileInfoCert对象，并按照key-value的形式存入couchdb
// =========================================================================================
func (t *CourtFileCertChaincode) AddRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//校验参数个数是否正确
	if len(args) != 4 {
		return shim.Error(AddRecordParameterErrStr + string(len(args)))
	}
	//从入参中获取所有必要的参数
	externalId := strings.TrimSpace(args[0])      //第一个参数bizId
	fileHash := strings.TrimSpace(args[1])        //第二个参数fileHash
	creationDateStr := strings.TrimSpace(args[2]) //第三个参数时间戳字符串
	metadataJson := strings.TrimSpace(args[3])    //第四个参数元数据json

	//对时间戳进行校验
	creationDate, err := strconv.ParseInt(creationDateStr, 10, 64)
	//必须是数字
	if err != nil {
		return shim.Error(StringToInt64ErrStr + err.Error())
	}
	//时间戳不能小于等于0
	if creationDate <= 0 {
		return shim.Error(InvalidCreationDateStr + err.Error())
	}
	//交易被创建时客户端的时间戳，从交易的ChannelHeader中提取
	txTime, err := stub.GetTxTimestamp()
	//获取时间戳异常处理
	if err != nil {
		return shim.Error(TxTimestampErrStr + err.Error())
	}
	txTimeInt := int64(txTime.GetSeconds())       //将时间戳类型转为int64
	txTimeStr := strconv.FormatInt(txTimeInt, 10) //将int64转为字符串
	//将metadataJson转为Metadata结构体
	var metadata Metadata                                //定义Metadata结构体类型常量
	e := json.Unmarshal([]byte(metadataJson), &metadata) //将传进来的json元数据转为Metadata类型
	//Json转结构体异常处理
	if e != nil {
		return shim.Error(UnmarshalMetadataErrStr + err.Error())
	}
	//交易提案时指定的交易ID
	txHash := stub.GetTxID()
	//定义常量keyId
	var bizId string
	//根据PreFileKey类型生成不同的keyId
	if metadata.ParentBizId == "" {
		//如果PreFileKey为空，keyId格式为"fileHash-txHash"
		bizId = fileHash + "-" + txHash
		//构建courtFileInfoCert对象
		courtFileInfoCert := CourtFileInfoCert{bizId, externalId, fileHash, creationDate, metadata}
		//构建courtFileInfoCert对象异常处理
		if &courtFileInfoCert == nil {
			return shim.Error(CreateCourtFileInfoCertErrStr)
		}
		//存证初始状态错误处理
		if courtFileInfoCert.Metadata.Status != CourtFileInfoCertUnArchived {
			return shim.Error(FileStatusErrStr)
		}
		//将courtFileInfoCert序列化为json
		courtFileInfoCertJSONAsBytes, err := json.Marshal(courtFileInfoCert)
		//将courtFileInfoCert序列化为json异常处理
		if err != nil {
			return shim.Error(MarshalCourtFileInfoCertErrStr + err.Error())
		}
		//将courtFileInfoCert以key-value的形式存入区块链
		err = stub.PutState(bizId, courtFileInfoCertJSONAsBytes)
		//上链异常处理
		if err != nil {
			return shim.Error(PutStateErrStr + err.Error())
		}
	} else {
		//如果PreFileKey不为空，截取originalFileHash和originalTxHash
		split := strings.Split(metadata.ParentBizId, "-")
		//校验PreFileKey格式是否正确
		if len(split) < 2 {
			return shim.Error(PreFileKeyFormatErrStr)
		}
		//截取的第一个字符串为originalFileHash
		originalFileHash := split[0]
		//街区的第二个字符串为originalTxHash
		originalTxHash := split[1]
		//keyId格式为"originalFileHash-originalTxHash-fileHash-txHash"
		bizId = originalFileHash + "-" + originalTxHash + "-" + fileHash + "-" + txHash
		//构建courtFileInfoCert对象
		courtFileInfoCert := CourtFileInfoCert{bizId, externalId, fileHash, creationDate, metadata}
		//构建courtFileInfoCert对象异常处理
		if &courtFileInfoCert == nil {
			return shim.Error(CreateCourtFileInfoCertErrStr)
		}
		//存证初始状态错误处理
		if courtFileInfoCert.Metadata.Status != CourtFileInfoCertUnArchived {
			return shim.Error(FileStatusErrStr)
		}
		//序列化courtFileInfoCert
		courtFileInfoCertJSONAsBytes, err := json.Marshal(courtFileInfoCert)
		//序列化courtFileInfoCert异常处理
		if err != nil {
			return shim.Error(MarshalCourtFileInfoCertErrStr + err.Error())
		}
		//存入couchdb异常处理
		e := stub.PutState(bizId, courtFileInfoCertJSONAsBytes)
		if e != nil {
			return shim.Error(PutStateErrStr + e.Error())
		}
		//创建keyId的CompositeKey
		keyIdNameIndexKey, err := stub.CreateCompositeKey(compositekeyIdIndexName, []string{originalFileHash, originalTxHash, fileHash, txHash})
		//创建CompositeKey的异常处理
		if err != nil {
			return shim.Error(CreateCompositeKeyErrStr + err.Error())
		}
		//将创建keyId的CompositeKey存入couchdb数据库
		errkeyIdNameIndexKey := stub.PutState(keyIdNameIndexKey, []byte{0x00})
		////存入couchdb异常处理
		if errkeyIdNameIndexKey != nil {
			return shim.Error(PutStateErrStr + errkeyIdNameIndexKey.Error())
		}
		//为修改之前的文件添加WaterMark事件记录
		t.AddEvent(stub, []string{metadata.ParentBizId, "WaterMark", metadata.Uploader, txTimeStr})
	}
	//为相关存证作添加AddRecord事件记录
	t.AddEvent(stub, []string{bizId, "AddRecord", metadata.Uploader, txTimeStr})
	//返回存证在couchdb中的主键keyId
	return shim.Success([]byte(bizId))
}

// ===== GetRecord ========================================================
//  根据keyId获取存证信息
//	@param	bizId operator
//	@return bool
//	根据keyId从couchdb中获取存证信息并返回
// =========================================================================================
func (t *CourtFileCertChaincode) GetRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//校验参数
	if len(args) != 2 {
		return shim.Error(GetRecordParameterErrStr + string(len(args)))
	}
	bizId := args[0]    //第一个参数keyId
	operator := args[1] //第二个参数操作者
	//交易被创建时客户端的时间戳，从交易的ChannelHeader中提取
	txTime, err := stub.GetTxTimestamp()
	//获取时间戳异常处理
	if err != nil {
		return shim.Error(TxTimestampErrStr + err.Error())
	}
	//将时间戳转为int64
	txTimeInt := int64(txTime.GetSeconds())
	//将int64时间戳转为string
	txTimestamp := strconv.FormatInt(txTimeInt, 10)
	//根据keyId获取存证信息
	courtFileInfoCert, err := stub.GetState(bizId)
	//获取存证信息异常处理
	if err != nil {
		return shim.Error(GetStateErrStr)
	} else if courtFileInfoCert == nil {
		return shim.Error(courtFileInfoCertNotFoundStr)
	}
	//为相关存证添加GetRecord事件记录
	t.AddEvent(stub, []string{bizId, "GetRecord", operator, txTimestamp})
	//返回存证信息
	return shim.Success(courtFileInfoCert)
}

// ===== GetAttestation(校验fileHash) ========================================================
//	@param	fileHash bizId operator
//	@return bool
//	根据keyId从couchdb中获取存证信息
//	将获取到的文件hash和参数传入的文件哈希进行比较
//	如果两者一致，返回true，否则返回true
// =========================================================================================
func (t *CourtFileCertChaincode) GetAttestation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 3 {
		return shim.Error(GetAttestationParameterErrStr + string(len(args)))
	}
	fileHash := args[0] //第一个参数要校验的文件哈希
	bizId := args[1]    //第二个参数存证的keyId
	operator := args[2] //第三个参数操作者
	//交易被创建时客户端的时间戳，从交易的ChannelHeader中提取
	txTime, err := stub.GetTxTimestamp()
	//获取时间戳异常处理
	if err != nil {
		return shim.Error(TxTimestampErrStr + err.Error())
	}
	txTimeInt := int64(txTime.GetSeconds())         //将时间戳转为int64
	txTimestamp := strconv.FormatInt(txTimeInt, 10) //将int64转为string
	//根据keyId获取存证信息
	courtFileInfoCertBytes, err := stub.GetState(bizId)
	//获取存证信息异常处理
	if err != nil {
		return shim.Error(GetStateErrStr + err.Error())
	}
	//创建courtFileInfoCert变量并取址
	courtFileInfoCert := &CourtFileInfoCert{}
	//将courtFileInfoCertBytes转为CourtFileInfoCert结构体
	json.Unmarshal([]byte(courtFileInfoCertBytes), &courtFileInfoCert)
	//将传入的文件哈希和获取到的文件哈希做对比,如果一致返回true
	if courtFileInfoCert.FileHash == fileHash {
		return shim.Success([]byte("true"))
	}
	//为相关文件添加GetAttestation事件记录
	t.AddEvent(stub, []string{bizId, "GetAttestation", operator, string(txTimestamp)})
	//如果获取的文件哈希与传入的文件哈希不一致返回false
	return shim.Success([]byte("false"))
}

// TODO 写ReadMe文档 写测试 告诉 foy
// ===== AttestationMetaDataHash(校验metaHash) ========================================================
//	@param	fileHash bizId operator
//	@return bool
//	根据keyId从couchdb中获取存证信息
//	将获取到的文件metaDataHash和参数传入的metaDataHash进行比较
//	如果两者一致，返回true，否则返回true
// =========================================================================================
func (t *CourtFileCertChaincode) AttestationMetaDataHash(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 3 {
		return shim.Error(GetAttestationParameterErrStr + string(len(args)))
	}
	metaDataHash := args[0] //第一个参数文件的元数据哈希
	bizId := args[1]        //第二个参数存证的keyId
	operator := args[2]     //第三个参数操作者
	//交易被创建时客户端的时间戳，从交易的ChannelHeader中提取
	txTime, err := stub.GetTxTimestamp()
	//获取时间戳异常处理
	if err != nil {
		return shim.Error(TxTimestampErrStr + err.Error())
	}
	txTimeInt := int64(txTime.GetSeconds())             //将时间戳转为int64
	txTimestamp := strconv.FormatInt(txTimeInt, 10)     //将int64转为string
	courtFileInfoCertBytes, err := stub.GetState(bizId) //根据keyId获取存证信息
	//取存证信息异常处理
	if err != nil {
		return shim.Error(GetStateErrStr + err.Error())
	}
	//创建CourtFileInfoCert变量并取址
	courtFileInfoCert := &CourtFileInfoCert{}
	//将courtFileInfoCertBytes转为CourtFileInfoCert结构体
	json.Unmarshal([]byte(courtFileInfoCertBytes), &courtFileInfoCert)
	//对比获取到的MetadataHash和传入的MetadataHash，如果一致返回true
	if courtFileInfoCert.Metadata.MetadataHash == metaDataHash {
		return shim.Success([]byte("true"))
	}
	//为相关文件添加GetAttestation事件记录
	t.AddEvent(stub, []string{bizId, "AttestationMetaDataHash", operator, string(txTimestamp)})
	//如果获取的MetadataHash与传入的MetadataHash不一致返回false
	return shim.Success([]byte("false"))
}

// ===== AddEvent ========================================================
//	为某个存证添加事件记录信息，这样我们可以清晰的看到对这个文件的一系列操作
//	@param	bizId eventName operator operator timestampStr
//	@return txHash
//	按照Event-bizId-eventName-operator-timestamp的格式为事件创建compositeKey
//  并存入couchdb数据库
//	方便后期对Event进行查询
// =========================================================================================
func (t *CourtFileCertChaincode) AddEvent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 4 {
		return shim.Error(AddEventParameterErrStr + string(len(args)))
	}
	bizId := args[0]        //第一个参数keyId
	eventName := args[1]    //第二个参数事件名称
	operator := args[2]     //第三个参数操作者
	timestampStr := args[3] //第四个参数时间戳
	//交易提案时指定的交易ID
	txHash := stub.GetTxID()
	//对timestampStr进行校验
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	//时间戳必须为数字字符串，否则返回Error
	if err != nil {
		return shim.Error(timeStampStrToInt64ErrStr + err.Error())
	}
	//timestamp不能为负数
	if timestamp <= 0 {
		return shim.Error(InvalidTimeStampStr)
	}
	//按照"Event~bizId~eventName~operator~timestamp"格式创建CompositeKey
	eventNameIndexKey, err := stub.CreateCompositeKey(compositeIndexName, []string{"Event", bizId, eventName, operator, timestampStr})
	//创建CompositeKey异常处理
	if err != nil {
		return shim.Error(CreateCompositeKeyErrStr + err.Error())
	}
	//将创建的CompositeKey存入couchdb
	errCompositeKey := stub.PutState(eventNameIndexKey, []byte{0x00})
	//CompositeKey存入couchdb异常处理
	if errCompositeKey != nil {
		return shim.Error(PutStateErrStr)
	}
	//返回txHash（交易提案时指定的交易ID）
	return shim.Success([]byte(txHash))
}

// ===== SearchEvent ========================================================
//  根据keyId查询这个存证的一系列事件记录，并分页返回
//	@param	bizId	pageSize	bookmark
//	@return SearchEventRes
//	根据传入的参数查询相应的SearchEventRes
// =========================================================================================
func (t *CourtFileCertChaincode) SearchEvent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 3 {
		return shim.Error(SearchEventParameterErrStr + string(len(args)))
	}
	//第一个参数keyId
	bizId := args[0]
	//第二个参数pageSize，用于指定每次查询几条记录
	pageSize, err := strconv.ParseInt(args[1], 10, 32) //第二个参数 pageSize
	//对pageSize进行校验
	if err != nil {
		return shim.Error(PageSizeParaFormatErrStr + err.Error())
	}
	//pageSize的值不能为0和负数，不能大于100，重置pageSize为10 ，查询10条记录
	if pageSize <= 0 || pageSize >= 100 {
		pageSize = 10
	}
	//第三个参数bookmark，用于指定从哪里开始查询
	bookmark := args[2]
	//根据CompositeKey分页查询
	//eventResults 			查询结果
	//eventResultsMetadata  分页信息
	eventResults, eventResultsMetadata, err := stub.GetStateByPartialCompositeKeyWithPagination(compositeIndexName, []string{"Event", bizId}, int32(pageSize), bookmark)
	//查询异常处理
	if err != nil {
		return shim.Error(GetStateByPartialCompositeKeyErrStr + err.Error())
	}
	//判断eventResults是否为空
	if eventResults != nil {
		defer eventResults.Close()
		//构建fileEventArr
		var fileEventsArr []FileEvent
		//对查询结果进行遍历
		for eventResults.HasNext() {
			//取出eventResults中的元素
			event, err := eventResults.Next()
			//异常处理
			if err != nil {
				return shim.Error(EventResultsNextErrStr + err.Error())
			}
			//对元素的key进行分割
			_, compositeKeyParts, err := stub.SplitCompositeKey(event.Key)
			//分割异常处理
			if err != nil {
				return shim.Error(SplitCompositeKeyErrStr + err.Error())
			}
			//获取Event必要字段
			Event := compositeKeyParts[0]
			bizId := compositeKeyParts[1]
			eventName := compositeKeyParts[2]
			operator := compositeKeyParts[3]
			timeStamp := compositeKeyParts[4]
			fileEvent := FileEvent{Event, bizId, eventName, operator, timeStamp}
			fileEventsArr = append(fileEventsArr, fileEvent)
		}
		//定义SearchEventRes变量，并赋值
		var searchRes SearchEventRes
		searchRes.FileEvents = fileEventsArr
		searchRes.ResponseMetadata.Bookmark = eventResultsMetadata.Bookmark
		searchRes.ResponseMetadata.RecordsCount = eventResultsMetadata.FetchedRecordsCount
		//将SearchEventRes序列化
		searchEventsByte, err := json.Marshal(searchRes)
		//序列化异常处理
		if err != nil {
			return shim.Error(MarshalsearchResErrStr + err.Error())
		}
		return shim.Success(searchEventsByte)
	}
	return shim.Success([]byte(""))
}

// ===== Archive ========================================================
//  如果存证对应的case已完成，将对应的状态改为已归档
//	@param	bizId	operator
//	@return txID
//	根据传入的keyId获取存证信息，将Status改为1，重新存入couchdb
// =========================================================================================
func (t *CourtFileCertChaincode) Archive(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 2 {
		return shim.Error(ArchiveParameterErrStr + string(len(args)))
	}
	//第一个参数keyId
	bizId := args[0]
	//第二个参数操作者
	operator := args[1]
	//交易被创建时客户端的时间戳，从交易的ChannelHeader中提取
	txTime, err := stub.GetTxTimestamp()
	//获取时间戳异常处理
	if err != nil {
		return shim.Error(TxTimestampErrStr + err.Error())
	}
	txTimeInt := int64(txTime.GetSeconds())         //将时间戳转为int64
	txTimestamp := strconv.FormatInt(txTimeInt, 10) //将int64转为string
	//根据keyId获取存证信息
	courtFileInfoCertBytes, err := stub.GetState(bizId)
	if err != nil {
		return shim.Error(GetStateErrStr + err.Error())
	}
	//定义courtFileInfoCert变量并取址
	courtFileInfoCert := &CourtFileInfoCert{}
	//将courtFileInfoCertBytes转为结构体
	json.Unmarshal([]byte(courtFileInfoCertBytes), &courtFileInfoCert)
	//将Status改为已归档
	courtFileInfoCert.Metadata.Status = CourtFileInfoCertArchived
	//序列化
	newCourtFileInfoCertBytes, e := json.Marshal(courtFileInfoCert)
	//序列化异常处理
	e = stub.PutState(bizId, newCourtFileInfoCertBytes)
	if e != nil {
		return shim.Error(bizId + PutStateErrStr + "," + e.Error())
	}
	//交易提案时指定的交易ID
	txID := stub.GetTxID()
	//添加Archive事件记录信息
	t.AddEvent(stub, []string{bizId, "Archive", operator, string(txTimestamp)})
	//返回txID
	return shim.Success([]byte(txID))
}

// ===== generalSearch ========================================================
//  根据queryString进行分页富查询
//	@param	queryString	pageSize bookmark
//	@return searchRes
//	根据queryString进行复查询，返回查询结果
// =========================================================================================
func (t *CourtFileCertChaincode) Search(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 3 {
		return shim.Error(SearchParameterErrStr + string(len(args)))
	}
	//第一个参数 根据需要拼出的queryString
	//例：queryString := fmt.Sprintf("{\"selector\":{\"$and\":[{\"externalId\":\"%s\"},{\"fileUri\":\"%s\"}]}}", externalId, fileUri)
	queryString := args[0]
	pageSize, err := strconv.ParseInt(args[1], 10, 32) //第二个参数 pageSize
	//对pageSize进行校验
	if err != nil {
		return shim.Error(PageSizeParaFormatErrStr + err.Error())
	}
	//pageSize的值不能为0和负数，不能大于100，重置pageSize为10 ，查询10条记录
	if pageSize <= 0 || pageSize >= 100 {
		pageSize = 10
	}
	bookmark := args[2] //第三个参数 pageSize
	//用queryString进行查询
	queryResults, queryMetadata, err := stub.GetQueryResultWithPagination(queryString, int32(pageSize), bookmark)
	//查询异常处理
	if err != nil {
		return shim.Error(QueryErrStr + err.Error())
	}
	//对查询结果进行校验
	if queryResults != nil {
		defer queryResults.Close()
		var courtFileInfoCertArr []CourtFileInfoCert
		var newFile CourtFileInfoCert
		//遍历查询结果
		for queryResults.HasNext() {
			file, err := queryResults.Next()
			if err != nil {
				return shim.Error(QueryResultsErrStr + err.Error())
			}
			//将结果中的元素转为CourtFileInfoCert结构体
			json.Unmarshal([]byte(file.Value), &newFile)
			json.Marshal(newFile)
			//将遍历的元素加入courtFileInfoCertArr数组
			courtFileInfoCertArr = append(courtFileInfoCertArr, newFile)
		}
		//创建searchRes变量并赋值
		var searchRes SearchRes
		searchRes.CourtFileInfoCerts = courtFileInfoCertArr
		searchRes.ResponseMetadata.Bookmark = queryMetadata.Bookmark
		searchRes.ResponseMetadata.RecordsCount = queryMetadata.FetchedRecordsCount
		//序列化
		searchResByte, err := json.Marshal(searchRes)
		//序列化异常处理
		if err != nil {
			return shim.Error(MarshalsearchResErrStr + err.Error())
		}
		//将查询结果返回
		return shim.Success(searchResByte)
	}
	return shim.Success([]byte(""))
}

// ===== OriginalFileKeyIdSearch ========================================================
//	查询所有包含OriginalFileHash的存证
//	@param	originalFileKeyId
//	@return searchRes
//	定义一个CourtFileInfoCert类型数组
//  首先将OriginalFileKeyId对应的存证信息加入数组
//  然后按照compositeKey查询所有与originalFileKeyId相关的存证
// =========================================================================================
func (t *CourtFileCertChaincode) OriginalFileKeyIdSearch(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//参数校验
	if len(args) != 1 {
		return shim.Error(OriginalFileKeyIdSearchParameterErrStr + string(len(args)))
	}
	//传入原始存证keyId
	originalFileKeyId := args[0]
	var originalFile CourtFileInfoCert
	var newFile CourtFileInfoCert
	//创建CourtFileInfoCert数组
	var courtFileInfoCertArr []CourtFileInfoCert
	//获取原始存证
	originalFileByte, err := stub.GetState(originalFileKeyId)
	//获取存证异常处理
	if err != nil {
		return shim.Error(GetOriginalFileErrStr + originalFileKeyId + "," + err.Error())
	}
	//将originalFileByte转为结构体
	json.Unmarshal(originalFileByte, &originalFile)
	//将原始存证加入CourtFileInfoCert数组
	courtFileInfoCertArr = append(courtFileInfoCertArr, originalFile)
	//对originalFileKeyId进行分割
	originalFileKeyIdSplit := strings.Split(originalFileKeyId, "-")
	//originalFileKeyId格式校验
	if len(originalFileKeyIdSplit) != 2 {
		return shim.Error(OriginalFileKeyIdSplitLen)
	}
	//组装compositekey查询的必要字段
	originalFileHash := originalFileKeyIdSplit[0]
	originalTxHash := originalFileKeyIdSplit[1]
	//根据compositekey进行查询
	originalFileResults, err := stub.GetStateByPartialCompositeKey(compositekeyIdIndexName, []string{originalFileHash, originalTxHash})
	//查询异常处理
	if err != nil {
		return shim.Error(GetStateByPartialCompositeKeyErrStr + err.Error())
	}
	defer originalFileResults.Close()
	//对查询结构进行遍历
	for originalFileResults.HasNext() {
		file, err := originalFileResults.Next()
		//遍历异常处理
		if err != nil {
			return shim.Error(originalFileResultsNextErrStr + err.Error())
		}
		//分割file.Key
		_, compositeKeyParts, err := stub.SplitCompositeKey(file.Key)
		if err != nil {
			return shim.Error(SplitCompositeKeyErrStr + err.Error())
		}
		//file.Key校验
		if len(compositeKeyParts) < 4 {
			return shim.Error(InvalidCompositeKey)
		}
		//组装keyId
		bizId := compositeKeyParts[0] + "-" + compositeKeyParts[1] + "-" + compositeKeyParts[2] + "-" + compositeKeyParts[3]
		//根据keyId获取存证信息
		fileByte, err := stub.GetState(bizId)
		//获取存证信息异常处理
		if err != nil {
			return shim.Error(GetStateErrStr + file.Key + "," + err.Error())
		}
		//将fileByte转为结构体
		json.Unmarshal(fileByte, &newFile)
		//将newFile加入courtFileInfoCertArr数组
		courtFileInfoCertArr = append(courtFileInfoCertArr, newFile)
	}
	//将courtFileInfoCertArr数组序列化
	courtFilesByte, err := json.Marshal(courtFileInfoCertArr)
	//序列化异常处理
	if err != nil {
		return shim.Error(MarshalCourtFileInfoCertErrStr)
	}
	//返回查询结果
	return shim.Success(courtFilesByte)
}

// TODO 打水印
// TODO  跟分布式结合存储
