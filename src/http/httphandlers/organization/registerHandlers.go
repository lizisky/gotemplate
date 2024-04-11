package organization

import "lizisky.com/lizisky/src/http/httpServer"

// RegisterHandlers register http request handlers
func RegisterHandlers() {
	httpServer.RegisterHandler(&requestHandler_getOrgList{})
	httpServer.RegisterHandler(&requestHandler_getAllOrgCategories{})

	httpServer.RegisterHandler(&requestHandler_createOrganization{})
	httpServer.RegisterHandler(&requestHandler_readOrgInfo{})
	httpServer.RegisterHandler(&requestHandler_updateOrgInfo{})

	httpServer.RegisterHandler(&requestHandler_addOrgStaff{})
	httpServer.RegisterHandler(&requestHandler_readOrgStaff{})
	httpServer.RegisterHandler(&requestHandler_updateOrgStaff{})
	httpServer.RegisterHandler(&requestHandler_removeOrgStaff{})

	// httpServer.RegisterHandler(&requestHandler_applyJoinOrg{})
	// httpServer.RegisterHandler(&requestHandler_getMembershipList{})
	// httpServer.RegisterHandler(&requestHandler_updateMemberInfo{})
	// httpServer.RegisterHandler(&requestHandler_memberSignIn{})
	// httpServer.RegisterHandler(&requestHandler_getMemberSignInRecord{})
}
