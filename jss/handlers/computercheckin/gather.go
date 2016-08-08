package computercheckin

import (
	"encoding/xml"
	"github.com/mosen/putj/jss"
	"net/http"
	"fmt"
)

type ComputerCheckin struct {
	XMLName xml.Name `xml:"computer_check_in" json:"-"`
	CheckInFrequency int `xml:"check_in_frequency" json:"check_in_frequency"`  // Minutes
	CreateStartupScript bool `xml:"create_startup_script" json:"create_startup_script:"`
	LogStartupEvent bool `xml:"log_startup_event" json:"log_startup_event"`
	CheckForPoliciesAtStartup bool `xml:"check_for_policies_at_startup" json:"check_for_policies_at_startup"`
	ApplyComputerLevelManagedPreferences bool `xml:"apply_computer_level_managed_preferences" json:"apply_computer_level_managed_preferences"`
	EnsureSSHIsEnabled bool `xml:"ensure_ssh_is_enabled" json:"ensure_ssh_is_enabled"`
	CreateLoginLogoutHooks bool `xml:"create_login_logout_hooks" json:"create_login_logout_hooks"`
	LogUsername bool `xml:"log_username" json:"log_username"`
	CheckForPoliciesAtLoginLogout bool `xml:"check_for_policies_at_login_lockout" json:"check_for_policies_at_login_lockout"`
	ApplyUserLevelManagedPreferences bool `xml:"apply_user_level_managed_preferences" json:"apply_user_level_managed_preferences"`
	HideRestorePartition bool `xml:"hide_restore_partition" json:"hide_restore_partition"`
	PerformLoginActionsInBackground bool `xml:"perform_login_actions_in_background" json:"perform_login_actions_in_background"`
	DisplayStatusToUser bool `xml:"display_status_to_user" json:"display_status_to_user"`
}

func ComputerCheckinCaptureHandler(api *jss.Api, state map[string]interface{}) error {
	req, err := api.NewRequest("GET", "/JSSResource/computercheckin", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	checkin := &ComputerCheckin{}
	if err := xml.NewDecoder(res.Body).Decode(checkin); err != nil {
		return err
	}

	state["computer_check_in"] = checkin
	fmt.Printf("%v\n", checkin)

	return nil
}

func init() {
	jss.RegisterCaptureHandler("computer_check_in", ComputerCheckinCaptureHandler)
}


