// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"emq/queue"
	"flag"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  CreateQueueResponse createQueue(CreateQueueRequest request)")
	fmt.Fprintln(os.Stderr, "  void deleteQueue(DeleteQueueRequest request)")
	fmt.Fprintln(os.Stderr, "  void purgeQueue(PurgeQueueRequest request)")
	fmt.Fprintln(os.Stderr, "  SetQueueAttributesResponse setQueueAttribute(SetQueueAttributesRequest request)")
	fmt.Fprintln(os.Stderr, "  SetQueueQuotaResponse setQueueQuota(SetQueueQuotaRequest request)")
	fmt.Fprintln(os.Stderr, "  GetQueueInfoResponse getQueueInfo(GetQueueInfoRequest request)")
	fmt.Fprintln(os.Stderr, "  ListQueueResponse listQueue(ListQueueRequest request)")
	fmt.Fprintln(os.Stderr, "  SetQueueRedrivePolicyResponse setQueueRedrivePolicy(SetQueueRedrivePolicyRequest request)")
	fmt.Fprintln(os.Stderr, "  void removeQueueRedrivePolicy(RemoveQueueRedrivePolicyRequest request)")
	fmt.Fprintln(os.Stderr, "  void setPermission(SetPermissionRequest request)")
	fmt.Fprintln(os.Stderr, "  void revokePermission(RevokePermissionRequest request)")
	fmt.Fprintln(os.Stderr, "  QueryPermissionResponse queryPermission(QueryPermissionRequest request)")
	fmt.Fprintln(os.Stderr, "  QueryPermissionForIdResponse queryPermissionForId(QueryPermissionForIdRequest request)")
	fmt.Fprintln(os.Stderr, "  ListPermissionsResponse listPermissions(ListPermissionsRequest request)")
	fmt.Fprintln(os.Stderr, "  CreateTagResponse createTag(CreateTagRequest request)")
	fmt.Fprintln(os.Stderr, "  void deleteTag(DeleteTagRequest request)")
	fmt.Fprintln(os.Stderr, "  GetTagInfoResponse getTagInfo(GetTagInfoRequest request)")
	fmt.Fprintln(os.Stderr, "  ListTagResponse listTag(ListTagRequest request)")
	fmt.Fprintln(os.Stderr, "  TimeSeriesData queryMetric(QueryMetricRequest request)")
	fmt.Fprintln(os.Stderr, "  QueryPrivilegedQueueResponse queryPrivilegedQueue(QueryPrivilegedQueueRequest request)")
	fmt.Fprintln(os.Stderr, "  VerifyEMQAdminResponse verifyEMQAdmin()")
	fmt.Fprintln(os.Stderr, "  VerifyEMQAdminRoleResponse verifyEMQAdminRole(VerifyEMQAdminRoleRequest request)")
	fmt.Fprintln(os.Stderr, "  void copyQueue(CopyQueueRequest request)")
	fmt.Fprintln(os.Stderr, "  GetQueueMetaResponse getQueueMeta(string queueName)")
	fmt.Fprintln(os.Stderr, "  Version getServiceVersion()")
	fmt.Fprintln(os.Stderr, "  void validClientVersion(Version clientVersion)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := queue.NewQueueServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "createQueue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateQueue requires 1 args")
			flag.Usage()
		}
		arg76 := flag.Arg(1)
		mbTrans77 := thrift.NewTMemoryBufferLen(len(arg76))
		defer mbTrans77.Close()
		_, err78 := mbTrans77.WriteString(arg76)
		if err78 != nil {
			Usage()
			return
		}
		factory79 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt80 := factory79.GetProtocol(mbTrans77)
		argvalue0 := queue.NewCreateQueueRequest()
		err81 := argvalue0.Read(jsProt80)
		if err81 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CreateQueue(value0))
		fmt.Print("\n")
		break
	case "deleteQueue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteQueue requires 1 args")
			flag.Usage()
		}
		arg82 := flag.Arg(1)
		mbTrans83 := thrift.NewTMemoryBufferLen(len(arg82))
		defer mbTrans83.Close()
		_, err84 := mbTrans83.WriteString(arg82)
		if err84 != nil {
			Usage()
			return
		}
		factory85 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt86 := factory85.GetProtocol(mbTrans83)
		argvalue0 := queue.NewDeleteQueueRequest()
		err87 := argvalue0.Read(jsProt86)
		if err87 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteQueue(value0))
		fmt.Print("\n")
		break
	case "purgeQueue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PurgeQueue requires 1 args")
			flag.Usage()
		}
		arg88 := flag.Arg(1)
		mbTrans89 := thrift.NewTMemoryBufferLen(len(arg88))
		defer mbTrans89.Close()
		_, err90 := mbTrans89.WriteString(arg88)
		if err90 != nil {
			Usage()
			return
		}
		factory91 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt92 := factory91.GetProtocol(mbTrans89)
		argvalue0 := queue.NewPurgeQueueRequest()
		err93 := argvalue0.Read(jsProt92)
		if err93 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PurgeQueue(value0))
		fmt.Print("\n")
		break
	case "setQueueAttribute":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SetQueueAttribute requires 1 args")
			flag.Usage()
		}
		arg94 := flag.Arg(1)
		mbTrans95 := thrift.NewTMemoryBufferLen(len(arg94))
		defer mbTrans95.Close()
		_, err96 := mbTrans95.WriteString(arg94)
		if err96 != nil {
			Usage()
			return
		}
		factory97 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt98 := factory97.GetProtocol(mbTrans95)
		argvalue0 := queue.NewSetQueueAttributesRequest()
		err99 := argvalue0.Read(jsProt98)
		if err99 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.SetQueueAttribute(value0))
		fmt.Print("\n")
		break
	case "setQueueQuota":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SetQueueQuota requires 1 args")
			flag.Usage()
		}
		arg100 := flag.Arg(1)
		mbTrans101 := thrift.NewTMemoryBufferLen(len(arg100))
		defer mbTrans101.Close()
		_, err102 := mbTrans101.WriteString(arg100)
		if err102 != nil {
			Usage()
			return
		}
		factory103 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt104 := factory103.GetProtocol(mbTrans101)
		argvalue0 := queue.NewSetQueueQuotaRequest()
		err105 := argvalue0.Read(jsProt104)
		if err105 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.SetQueueQuota(value0))
		fmt.Print("\n")
		break
	case "getQueueInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetQueueInfo requires 1 args")
			flag.Usage()
		}
		arg106 := flag.Arg(1)
		mbTrans107 := thrift.NewTMemoryBufferLen(len(arg106))
		defer mbTrans107.Close()
		_, err108 := mbTrans107.WriteString(arg106)
		if err108 != nil {
			Usage()
			return
		}
		factory109 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt110 := factory109.GetProtocol(mbTrans107)
		argvalue0 := queue.NewGetQueueInfoRequest()
		err111 := argvalue0.Read(jsProt110)
		if err111 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetQueueInfo(value0))
		fmt.Print("\n")
		break
	case "listQueue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ListQueue requires 1 args")
			flag.Usage()
		}
		arg112 := flag.Arg(1)
		mbTrans113 := thrift.NewTMemoryBufferLen(len(arg112))
		defer mbTrans113.Close()
		_, err114 := mbTrans113.WriteString(arg112)
		if err114 != nil {
			Usage()
			return
		}
		factory115 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt116 := factory115.GetProtocol(mbTrans113)
		argvalue0 := queue.NewListQueueRequest()
		err117 := argvalue0.Read(jsProt116)
		if err117 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ListQueue(value0))
		fmt.Print("\n")
		break
	case "setQueueRedrivePolicy":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SetQueueRedrivePolicy requires 1 args")
			flag.Usage()
		}
		arg118 := flag.Arg(1)
		mbTrans119 := thrift.NewTMemoryBufferLen(len(arg118))
		defer mbTrans119.Close()
		_, err120 := mbTrans119.WriteString(arg118)
		if err120 != nil {
			Usage()
			return
		}
		factory121 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt122 := factory121.GetProtocol(mbTrans119)
		argvalue0 := queue.NewSetQueueRedrivePolicyRequest()
		err123 := argvalue0.Read(jsProt122)
		if err123 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.SetQueueRedrivePolicy(value0))
		fmt.Print("\n")
		break
	case "removeQueueRedrivePolicy":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RemoveQueueRedrivePolicy requires 1 args")
			flag.Usage()
		}
		arg124 := flag.Arg(1)
		mbTrans125 := thrift.NewTMemoryBufferLen(len(arg124))
		defer mbTrans125.Close()
		_, err126 := mbTrans125.WriteString(arg124)
		if err126 != nil {
			Usage()
			return
		}
		factory127 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt128 := factory127.GetProtocol(mbTrans125)
		argvalue0 := queue.NewRemoveQueueRedrivePolicyRequest()
		err129 := argvalue0.Read(jsProt128)
		if err129 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RemoveQueueRedrivePolicy(value0))
		fmt.Print("\n")
		break
	case "setPermission":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SetPermission requires 1 args")
			flag.Usage()
		}
		arg130 := flag.Arg(1)
		mbTrans131 := thrift.NewTMemoryBufferLen(len(arg130))
		defer mbTrans131.Close()
		_, err132 := mbTrans131.WriteString(arg130)
		if err132 != nil {
			Usage()
			return
		}
		factory133 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt134 := factory133.GetProtocol(mbTrans131)
		argvalue0 := queue.NewSetPermissionRequest()
		err135 := argvalue0.Read(jsProt134)
		if err135 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.SetPermission(value0))
		fmt.Print("\n")
		break
	case "revokePermission":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RevokePermission requires 1 args")
			flag.Usage()
		}
		arg136 := flag.Arg(1)
		mbTrans137 := thrift.NewTMemoryBufferLen(len(arg136))
		defer mbTrans137.Close()
		_, err138 := mbTrans137.WriteString(arg136)
		if err138 != nil {
			Usage()
			return
		}
		factory139 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt140 := factory139.GetProtocol(mbTrans137)
		argvalue0 := queue.NewRevokePermissionRequest()
		err141 := argvalue0.Read(jsProt140)
		if err141 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RevokePermission(value0))
		fmt.Print("\n")
		break
	case "queryPermission":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryPermission requires 1 args")
			flag.Usage()
		}
		arg142 := flag.Arg(1)
		mbTrans143 := thrift.NewTMemoryBufferLen(len(arg142))
		defer mbTrans143.Close()
		_, err144 := mbTrans143.WriteString(arg142)
		if err144 != nil {
			Usage()
			return
		}
		factory145 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt146 := factory145.GetProtocol(mbTrans143)
		argvalue0 := queue.NewQueryPermissionRequest()
		err147 := argvalue0.Read(jsProt146)
		if err147 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryPermission(value0))
		fmt.Print("\n")
		break
	case "queryPermissionForId":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryPermissionForId requires 1 args")
			flag.Usage()
		}
		arg148 := flag.Arg(1)
		mbTrans149 := thrift.NewTMemoryBufferLen(len(arg148))
		defer mbTrans149.Close()
		_, err150 := mbTrans149.WriteString(arg148)
		if err150 != nil {
			Usage()
			return
		}
		factory151 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt152 := factory151.GetProtocol(mbTrans149)
		argvalue0 := queue.NewQueryPermissionForIdRequest()
		err153 := argvalue0.Read(jsProt152)
		if err153 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryPermissionForId(value0))
		fmt.Print("\n")
		break
	case "listPermissions":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ListPermissions requires 1 args")
			flag.Usage()
		}
		arg154 := flag.Arg(1)
		mbTrans155 := thrift.NewTMemoryBufferLen(len(arg154))
		defer mbTrans155.Close()
		_, err156 := mbTrans155.WriteString(arg154)
		if err156 != nil {
			Usage()
			return
		}
		factory157 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt158 := factory157.GetProtocol(mbTrans155)
		argvalue0 := queue.NewListPermissionsRequest()
		err159 := argvalue0.Read(jsProt158)
		if err159 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ListPermissions(value0))
		fmt.Print("\n")
		break
	case "createTag":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateTag requires 1 args")
			flag.Usage()
		}
		arg160 := flag.Arg(1)
		mbTrans161 := thrift.NewTMemoryBufferLen(len(arg160))
		defer mbTrans161.Close()
		_, err162 := mbTrans161.WriteString(arg160)
		if err162 != nil {
			Usage()
			return
		}
		factory163 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt164 := factory163.GetProtocol(mbTrans161)
		argvalue0 := queue.NewCreateTagRequest()
		err165 := argvalue0.Read(jsProt164)
		if err165 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CreateTag(value0))
		fmt.Print("\n")
		break
	case "deleteTag":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteTag requires 1 args")
			flag.Usage()
		}
		arg166 := flag.Arg(1)
		mbTrans167 := thrift.NewTMemoryBufferLen(len(arg166))
		defer mbTrans167.Close()
		_, err168 := mbTrans167.WriteString(arg166)
		if err168 != nil {
			Usage()
			return
		}
		factory169 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt170 := factory169.GetProtocol(mbTrans167)
		argvalue0 := queue.NewDeleteTagRequest()
		err171 := argvalue0.Read(jsProt170)
		if err171 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteTag(value0))
		fmt.Print("\n")
		break
	case "getTagInfo":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTagInfo requires 1 args")
			flag.Usage()
		}
		arg172 := flag.Arg(1)
		mbTrans173 := thrift.NewTMemoryBufferLen(len(arg172))
		defer mbTrans173.Close()
		_, err174 := mbTrans173.WriteString(arg172)
		if err174 != nil {
			Usage()
			return
		}
		factory175 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt176 := factory175.GetProtocol(mbTrans173)
		argvalue0 := queue.NewGetTagInfoRequest()
		err177 := argvalue0.Read(jsProt176)
		if err177 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTagInfo(value0))
		fmt.Print("\n")
		break
	case "listTag":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ListTag requires 1 args")
			flag.Usage()
		}
		arg178 := flag.Arg(1)
		mbTrans179 := thrift.NewTMemoryBufferLen(len(arg178))
		defer mbTrans179.Close()
		_, err180 := mbTrans179.WriteString(arg178)
		if err180 != nil {
			Usage()
			return
		}
		factory181 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt182 := factory181.GetProtocol(mbTrans179)
		argvalue0 := queue.NewListTagRequest()
		err183 := argvalue0.Read(jsProt182)
		if err183 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ListTag(value0))
		fmt.Print("\n")
		break
	case "queryMetric":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryMetric requires 1 args")
			flag.Usage()
		}
		arg184 := flag.Arg(1)
		mbTrans185 := thrift.NewTMemoryBufferLen(len(arg184))
		defer mbTrans185.Close()
		_, err186 := mbTrans185.WriteString(arg184)
		if err186 != nil {
			Usage()
			return
		}
		factory187 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt188 := factory187.GetProtocol(mbTrans185)
		argvalue0 := queue.NewQueryMetricRequest()
		err189 := argvalue0.Read(jsProt188)
		if err189 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryMetric(value0))
		fmt.Print("\n")
		break
	case "queryPrivilegedQueue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryPrivilegedQueue requires 1 args")
			flag.Usage()
		}
		arg190 := flag.Arg(1)
		mbTrans191 := thrift.NewTMemoryBufferLen(len(arg190))
		defer mbTrans191.Close()
		_, err192 := mbTrans191.WriteString(arg190)
		if err192 != nil {
			Usage()
			return
		}
		factory193 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt194 := factory193.GetProtocol(mbTrans191)
		argvalue0 := queue.NewQueryPrivilegedQueueRequest()
		err195 := argvalue0.Read(jsProt194)
		if err195 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryPrivilegedQueue(value0))
		fmt.Print("\n")
		break
	case "verifyEMQAdmin":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "VerifyEMQAdmin requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.VerifyEMQAdmin())
		fmt.Print("\n")
		break
	case "verifyEMQAdminRole":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "VerifyEMQAdminRole requires 1 args")
			flag.Usage()
		}
		arg196 := flag.Arg(1)
		mbTrans197 := thrift.NewTMemoryBufferLen(len(arg196))
		defer mbTrans197.Close()
		_, err198 := mbTrans197.WriteString(arg196)
		if err198 != nil {
			Usage()
			return
		}
		factory199 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt200 := factory199.GetProtocol(mbTrans197)
		argvalue0 := queue.NewVerifyEMQAdminRoleRequest()
		err201 := argvalue0.Read(jsProt200)
		if err201 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.VerifyEMQAdminRole(value0))
		fmt.Print("\n")
		break
	case "copyQueue":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CopyQueue requires 1 args")
			flag.Usage()
		}
		arg202 := flag.Arg(1)
		mbTrans203 := thrift.NewTMemoryBufferLen(len(arg202))
		defer mbTrans203.Close()
		_, err204 := mbTrans203.WriteString(arg202)
		if err204 != nil {
			Usage()
			return
		}
		factory205 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt206 := factory205.GetProtocol(mbTrans203)
		argvalue0 := queue.NewCopyQueueRequest()
		err207 := argvalue0.Read(jsProt206)
		if err207 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CopyQueue(value0))
		fmt.Print("\n")
		break
	case "getQueueMeta":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetQueueMeta requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetQueueMeta(value0))
		fmt.Print("\n")
		break
	case "getServiceVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetServiceVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetServiceVersion())
		fmt.Print("\n")
		break
	case "validClientVersion":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ValidClientVersion requires 1 args")
			flag.Usage()
		}
		arg209 := flag.Arg(1)
		mbTrans210 := thrift.NewTMemoryBufferLen(len(arg209))
		defer mbTrans210.Close()
		_, err211 := mbTrans210.WriteString(arg209)
		if err211 != nil {
			Usage()
			return
		}
		factory212 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt213 := factory212.GetProtocol(mbTrans210)
		argvalue0 := queue.NewVersion()
		err214 := argvalue0.Read(jsProt213)
		if err214 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ValidClientVersion(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
