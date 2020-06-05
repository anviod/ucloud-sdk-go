// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/services/vpc"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/driver"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/utils"
	"github.com/ucloud/ucloud-sdk-go/ucloud/utest/validation"
)

func TestScenario2468(t *testing.T) {
	spec.ParallelTest(t, &driver.Scenario{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Id: "2468",
		Vars: func(scenario *driver.Scenario) map[string]interface{} {
			return map[string]interface{}{
				"acl_name":        "acl_auto_test1",
				"acl_name_modify": "acl_auto_test2",
				"priority1":       1,
				"IpProtocol1":     "TCP",
				"priority2":       100,
				"IpProtocol2":     "UDP",
				"Region":          "cn-bj2",
				"Zone":            "cn-bj2-02",
			}
		},
		Owners: []string{"li.wei@ucloud.cn"},
		Title:  "ACL自动化回归-基本操作",
		Steps: []*driver.Step{
			testStep2468DescribeImage01,
			testStep2468CreateVPC02,
			testStep2468CreateSubnet03,
			testStep2468CreateUHostInstance04,
			testStep2468CreateNetworkAcl05,
			testStep2468DescribeNetworkAcl06,
			testStep2468UpdateNetworkAcl07,
			testStep2468CreateNetworkAclAssociation08,
			testStep2468DescribeNetworkAclAssociation09,
			testStep2468DescribeNetworkAclAssociationBySubnet10,
			testStep2468GetNetworkAclTargetResource11,
			testStep2468CreateNetworkAclEntry12,
			testStep2468CreateNetworkAclEntry13,
			testStep2468DescribeNetworkAclEntry14,
			testStep2468UpdateNetworkAclEntry15,
			testStep2468DeleteNetworkAclEntry16,
			testStep2468DeleteNetworkAclEntry17,
			testStep2468DeleteNetworkAclAssociation18,
			testStep2468DeleteNetworkAcl19,
			testStep2468PoweroffUHostInstance20,
			testStep2468TerminateUHostInstance21,
			testStep2468DeleteSubnet22,
			testStep2468DeleteVPC23,
		},
	})
}

var testStep2468DescribeImage01 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewDescribeImageRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":      step.Scenario.GetVar("Zone"),
			"Region":    step.Scenario.GetVar("Region"),
			"OsType":    "Linux",
			"ImageType": "Base",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeImage(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("Image_Id", step.Must(utils.GetValue(resp, "ImageSet.0.ImageId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "DescribeImageResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取镜像列表",
	FastFail:      false,
}

var testStep2468CreateVPC02 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"Network": []interface{}{
				"192.168.0.0/16",
			},
			"Name": "acl-test",
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateVPC(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("vpc_id", step.Must(utils.GetValue(resp, "VPCId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建VPC",
	FastFail:      false,
}

var testStep2468CreateSubnet03 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":      step.Scenario.GetVar("vpc_id"),
			"SubnetName": "acl-subnet-test",
			"Subnet":     "192.168.11.0",
			"Region":     step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateSubnet(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("subnet_id", step.Must(utils.GetValue(resp, "SubnetId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建子网",
	FastFail:      false,
}

var testStep2468CreateUHostInstance04 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewCreateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":        step.Scenario.GetVar("Zone"),
			"VPCId":       step.Scenario.GetVar("vpc_id"),
			"Tag":         "Default",
			"SubnetId":    step.Scenario.GetVar("subnet_id"),
			"Region":      step.Scenario.GetVar("Region"),
			"Password":    "VXFhNzg5VGVzdCFAIyQ7LA==",
			"Name":        "acl-test-uhost",
			"Memory":      1024,
			"MachineType": "N",
			"LoginMode":   "Password",
			"ImageId":     step.Scenario.GetVar("Image_Id"),
			"Disks": []map[string]interface{}{
				{
					"IsBoot": "True",
					"Size":   20,
					"Type":   "LOCAL_NORMAL",
				},
			},
			"CPU": 1,
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("uhost_id", step.Must(utils.GetValue(resp, "UHostIds.0")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Action", "CreateUHostInstanceResponse", "str_eq"),
		}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建云主机",
	FastFail:      false,
}

var testStep2468CreateNetworkAcl05 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateNetworkAclRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VpcId":   step.Scenario.GetVar("vpc_id"),
			"Region":  step.Scenario.GetVar("Region"),
			"AclName": step.Scenario.GetVar("acl_name"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateNetworkAcl(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("acl_id", step.Must(utils.GetValue(resp, "AclId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(120) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建网络ACL",
	FastFail:      false,
}

var testStep2468DescribeNetworkAcl06 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDescribeNetworkAclRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeNetworkAcl(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("AclList", 0, "len_gt"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取网络ACL",
	FastFail:      false,
}

var testStep2468UpdateNetworkAcl07 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewUpdateNetworkAclRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":      step.Scenario.GetVar("Region"),
			"Description": "acltest",
			"AclName":     step.Scenario.GetVar("acl_name_modify"),
			"AclId":       step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateNetworkAcl(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "更改ACL",
	FastFail:      false,
}

var testStep2468CreateNetworkAclAssociation08 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateNetworkAclAssociationRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetworkId": step.Scenario.GetVar("subnet_id"),
			"Region":       step.Scenario.GetVar("Region"),
			"AclId":        step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateNetworkAclAssociation(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建ACL的绑定关系",
	FastFail:      false,
}

var testStep2468DescribeNetworkAclAssociation09 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDescribeNetworkAclAssociationRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"AclId":  step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeNetworkAclAssociation(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("AssociationList", 0, "len_gt"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取网络ACL的绑定关系列表",
	FastFail:      false,
}

var testStep2468DescribeNetworkAclAssociationBySubnet10 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDescribeNetworkAclAssociationBySubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetworkId": step.Scenario.GetVar("subnet_id"),
			"Region":       step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeNetworkAclAssociationBySubnet(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("Association.AclId", step.Scenario.GetVar("acl_id"), "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取子网的ACL绑定信息",
	FastFail:      false,
}

var testStep2468GetNetworkAclTargetResource11 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewGetNetworkAclTargetResourceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetworkId": []interface{}{
				step.Scenario.GetVar("subnet_id"),
			},
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.GetNetworkAclTargetResource(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("TargetResourceList", 1, "len_ge"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取ACL规则应用目标列表",
	FastFail:      false,
}

var testStep2468CreateNetworkAclEntry12 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateNetworkAclEntryRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":      step.Scenario.GetVar("Region"),
			"Priority":    step.Scenario.GetVar("priority1"),
			"PortRange":   234,
			"IpProtocol":  step.Scenario.GetVar("IpProtocol1"),
			"EntryAction": "Accept",
			"Direction":   "Ingress",
			"CidrBlock":   "1.1.1.1/32",
			"AclId":       step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateNetworkAclEntry(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("EntryId", step.Must(utils.GetValue(resp, "EntryId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建Acl的规则",
	FastFail:      false,
}

var testStep2468CreateNetworkAclEntry13 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewCreateNetworkAclEntryRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":      step.Scenario.GetVar("Region"),
			"Priority":    step.Scenario.GetVar("priority1"),
			"PortRange":   234,
			"IpProtocol":  step.Scenario.GetVar("IpProtocol1"),
			"EntryAction": "Accept",
			"Direction":   "Egress",
			"CidrBlock":   "1.1.1.1/32",
			"AclId":       step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.CreateNetworkAclEntry(req)
		if err != nil {
			return resp, err
		}

		step.Scenario.SetVar("EntryId2", step.Must(utils.GetValue(resp, "EntryId")))
		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "创建Acl的规则",
	FastFail:      false,
}

var testStep2468DescribeNetworkAclEntry14 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDescribeNetworkAclEntryRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"AclId":  step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DescribeNetworkAclEntry(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
			validation.Builtins.NewValidator("EntryList", 0, "len_gt"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "获取ACL的规则信息",
	FastFail:      false,
}

var testStep2468UpdateNetworkAclEntry15 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewUpdateNetworkAclEntryRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":      step.Scenario.GetVar("Region"),
			"Priority":    step.Scenario.GetVar("priority2"),
			"PortRange":   2323,
			"IpProtocol":  step.Scenario.GetVar("IpProtocol2"),
			"EntryId":     step.Scenario.GetVar("EntryId"),
			"EntryAction": "Reject",
			"Direction":   "Ingress",
			"CidrBlock":   "2.2.2.2/32",
			"AclId":       step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.UpdateNetworkAclEntry(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "更新Acl的规则",
	FastFail:      false,
}

var testStep2468DeleteNetworkAclEntry16 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteNetworkAclEntryRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"EntryId": step.Scenario.GetVar("EntryId"),
			"AclId":   step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteNetworkAclEntry(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除ACL的规则",
	FastFail:      false,
}

var testStep2468DeleteNetworkAclEntry17 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteNetworkAclEntryRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region":  step.Scenario.GetVar("Region"),
			"EntryId": step.Scenario.GetVar("EntryId2"),
			"AclId":   step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteNetworkAclEntry(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除ACL的规则",
	FastFail:      false,
}

var testStep2468DeleteNetworkAclAssociation18 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteNetworkAclAssociationRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetworkId": step.Scenario.GetVar("subnet_id"),
			"Region":       step.Scenario.GetVar("Region"),
			"AclId":        step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteNetworkAclAssociation(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除网络ACL绑定关系",
	FastFail:      false,
}

var testStep2468DeleteNetworkAcl19 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteNetworkAclRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Region": step.Scenario.GetVar("Region"),
			"AclId":  step.Scenario.GetVar("acl_id"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteNetworkAcl(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{
			validation.Builtins.NewValidator("RetCode", 0, "str_eq"),
		}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除网络ACL",
	FastFail:      false,
}

var testStep2468PoweroffUHostInstance20 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewPoweroffUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.PoweroffUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(0) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "模拟主机掉电",
	FastFail:      false,
}

var testStep2468TerminateUHostInstance21 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("UHost")
		if err != nil {
			return nil, err
		}
		client := c.(*uhost.UHostClient)

		req := client.NewTerminateUHostInstanceRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"Zone":    step.Scenario.GetVar("Zone"),
			"UHostId": step.Scenario.GetVar("uhost_id"),
			"Region":  step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.TerminateUHostInstance(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(30) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除云主机",
	FastFail:      false,
}

var testStep2468DeleteSubnet22 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteSubnetRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"SubnetId": step.Scenario.GetVar("subnet_id"),
			"Region":   step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteSubnet(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(5) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除子网",
	FastFail:      false,
}

var testStep2468DeleteVPC23 = &driver.Step{
	Invoker: func(step *driver.Step) (interface{}, error) {
		c, err := step.LoadFixture("VPC")
		if err != nil {
			return nil, err
		}
		client := c.(*vpc.VPCClient)

		req := client.NewDeleteVPCRequest()
		err = utils.SetRequest(req, map[string]interface{}{
			"VPCId":  step.Scenario.GetVar("vpc_id"),
			"Region": step.Scenario.GetVar("Region"),
		})
		if err != nil {
			return nil, err
		}

		resp, err := client.DeleteVPC(req)
		if err != nil {
			return resp, err
		}

		return resp, nil
	},
	Validators: func(step *driver.Step) []driver.TestValidator {
		return []driver.TestValidator{}
	},
	StartupDelay:  time.Duration(10) * time.Second,
	MaxRetries:    3,
	RetryInterval: 1 * time.Second,
	Title:         "删除VPC",
	FastFail:      false,
}