package base_info

import (
	pb "Open_IM/pkg/proto/group"
	open_im_sdk "Open_IM/pkg/proto/sdk_ws"
)

type CommResp struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
type Id2Result struct {
	UserID string `json:"userID"`
	Result int32  `json:"result"`
}

type KickGroupMemberReq struct {
	GroupID          string   `json:"groupID" binding:"required"`
	KickedUserIDList []string `json:"kickedUserIDList" binding:"required"`
	Reason           string   `json:"reason"`
	OperationID      string   `json:"operationID" binding:"required"`
}
type KickGroupMemberResp struct {
	CommResp
	Data []*Id2Result `json:"data"`
}

type GetGroupMembersInfoReq struct {
	GroupID     string   `json:"groupID" binding:"required"`
	MemberList  []string `json:"memberList" binding:"required"`
	OperationID string   `json:"operationID" binding:"required"`
}
type GetGroupMembersInfoResp struct {
	CommResp
	Data []*open_im_sdk.GroupMemberFullInfo `json:"data"`
}

type InviteUserToGroupReq struct {
	GroupID           string   `json:"groupID" binding:"required"`
	InvitedUserIDList []string `json:"uidList" binding:"required"`
	Reason            string   `json:"reason"`
	OperationID       string   `json:"operationID" binding:"required"`
}
type InviteUserToGroupResp struct {
	CommResp
	Data []Id2Result `json:"data"`
}

type GetJoinedGroupListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}
type GetJoinedGroupListResp struct {
	CommResp
	Data []*open_im_sdk.GroupInfo `json:"data"`
}

type GetGroupMemberListReq struct {
	GroupID     string `json:"groupID"`
	Filter      int32  `json:"filter"`
	NextSeq     int32  `json:"nextSeq"`
	OperationID string `json:"operationID"`
}
type GetGroupMemberListResp struct {
	CommResp
	NextSeq int32                              `json:"nextSeq"`
	Data    []*open_im_sdk.GroupMemberFullInfo `json:"data"`
}

type GetGroupAllMemberReq struct {
	GroupID     string `json:"groupID"`
	OperationID string `json:"operationID"`
}
type GetGroupAllMemberResp struct {
	CommResp
	Data []*open_im_sdk.GroupMemberFullInfo `json:"data"`
}

type CreateGroupReq struct {
	MemberList   []*pb.GroupAddMemberInfo `json:"memberList"`
	GroupName    string                   `json:"groupName"`
	Introduction string                   `json:"introduction"`
	Notification string                   `json:"notification"`
	FaceUrl      string                   `json:"faceUrl"`
	OperationID  string                   `json:"operationID" binding:"required"`
	GroupType    int32                    `json:"groupType"`
	Ex           string                   `json:"ex"`
}
type CreateGroupResp struct {
	CommResp
	Data open_im_sdk.GroupInfo `json:"data"`
}

type GetGroupApplicationListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"` //my application
}
type GetGroupApplicationListResp struct {
	CommResp
	Data []*open_im_sdk.GroupRequest `json:"data"`
}

type GetGroupInfoReq struct {
	GroupIDList []string `json:"groupIDList" binding:"required"`
	OperationID string   `json:"operationID" binding:"required"`
}
type GetGroupInfoResp struct {
	CommResp
	Data []open_im_sdk.GroupInfo `json:"data"`
}

type ApplicationGroupResponseReq struct {
	OperationID  string `json:"groupIDList" binding:"required"`
	GroupID      string `json:"groupIDList" binding:"required"`
	FromUserID   string `json:"groupIDList" binding:"required"`
	HandledMsg   string `json:"groupIDList" binding:"required"`
	HandleResult int32  `json:"groupIDList" binding:"required"`
}
type ApplicationGroupResponseResp struct {
	CommResp
}

type JoinGroupReq struct {
	GroupID     string `json:"groupID"`
	ReqMessage  string `json:"reqMessage"`
	OperationID string `json:"operationID"`
}
type JoinGroupResp struct {
	CommResp
}

type QuitGroupReq struct {
	GroupID     string `json:"groupID" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}
type QuitGroupResp struct {
	CommResp
}

type SetGroupInfoReq struct {
	open_im_sdk.GroupInfo
	OperationID string `json:"operationID" binding:"required"`
}
type SetGroupInfoResp struct {
	CommResp
}

type TransferGroupOwnerReq struct {
	GroupID        string `json:"groupID" binding:"required"`
	OldOwnerUserID string `json:"oldOwnerUserID" binding:"required"`
	NewOwnerUserID string `json:"newOwnerUserID" binding:"required"`
	OperationID    string `json:"operationID" binding:"required"`
}