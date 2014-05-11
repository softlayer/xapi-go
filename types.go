package xapi

import (
	"time"
)

type VirtualMachines []string

type Session struct {
	Auth_user_name     string
	Auth_user_sid      string
	Is_local_superuser bool
	Last_active        time.Time
	Other_config       map[string]string
	Parent             string
	Pool               bool
	Rbac_permissions   []string
	Subject            string
	Tasks              []string
	This_host          string
	This_user          string
	Uuid               string
	Validation_time    time.Time
}

type VDI struct {
	Allow_caching        bool
	Allowed_operations   []string
	Crash_dumps          []string
	Current_operations   map[string]string
	Is_a_snapshot        bool
	Location             string
	Managed              bool
	Metadata_latest      bool
	Missing              bool
	Name_description     string
	Name_label           string
	On_boot              string
	Other_config         map[string]string
	Parent               string
	Physcial_utilisation int
	Read_only            bool
	Sharable             bool
	Sm_config            map[string]string
	Snapshot_of          []string
	SR                   []string
	Storage_lock         bool
	Tags                 string
	Type                 string
	Uuid                 string
	VDBs                 []string
	Virtual_size         int
	Xenstore_data        map[string]string
}

type VDB struct {
	Allowed_operations       []string
	Current_operations       map[string]string
	Bootable                 bool
	Currently_attached       bool
	Device                   string
	Empty                    bool
	Metrics                  string
	Mode                     string
	Other_config             map[string]string
	Qos_algorithm_params     map[string]string
	Qos_algorithm_type       string
	Qos_supported_algorithms []string
	Runtime_properties       map[string]string
	Status_code              int
	Status_detail            string
	Uuid                     string
	VDI                      string
	VM                       string
	Storage_lock             bool
	Type                     string
	Unpluggable              bool
	Userdevice               string
}

type VM struct {
	Actions_after_crash       string
	Actions_after_reboot      string
	Actions_after_shutdown    string
	Affinity                  string
	Allowed_operations        []string
	Appliance                 string
	Attached_PCIs             []string
	Bios_strings              map[string]string
	Blobs                     map[string][]byte
	Blocked_operations        map[string]string
	Children                  []string
	Consoles                  []string
	Crash_dumps               []string
	Current_operations        map[string]string
	Domarch                   string
	Domid                     int
	Guest_metrics             string
	Ha_always_run             bool
	Ha_restart_priority       string
	HVM_boot_params           map[string]string
	HVM_boot_policy           string
	HVM_shadow_multiplier     float64
	Is_a_snapshot             bool
	Is_a_template             bool
	Is_control_domain         bool
	Is_snapshot_from_vmpp     bool
	Last_boot_CPU_flags       map[string]string
	Last_booted_record        string
	Memory_dynamic_max        int
	Memory_dynamic_min        int
	Memory_overhead           int
	Memory_static_max         int
	Memory_static_min         int
	Memory_target             int
	Metrics                   string
	Name_description          string
	Name_label                string
	Order                     int
	Other_config              map[string]string
	Parent                    string
	PCI_bus                   string
	Platform                  map[string]string
	Power_status              string
	Protection_policy         string
	PV_args                   string
	PV_bootloader             string
	PV_bootloader_args        string
	PV_kernel                 string
	PV_legacy_args            string
	PV_ramdisk                string
	Recommendations           string
	Resident_on               string
	Shutdown_delay            int
	Snapshot_info             map[string]string
	Snapshot_metadata         string
	Snapshot_time             time.Time
	Start_delay               int
	Suspend_SR                string
	Suspend_VDI               string
	Tags                      []string
	Transportable_snapshot_id string
	User_version              int
	Uuid                      string
	VBDs                      []string
	VCPUs_at_startup          int
	VCPUs_max                 int
	VCPUs_params              map[string]string
	Version                   int
	VGPUs                     []string
	VIFs                      []string
	VTPM                      []string
	Xenstore_data             map[string]string
}

type Event struct {
	Class     string
	Id        int
	Obj_uuid  string
	Operation string
	Ref       string
	Datetime  time.Time
}

type VIF struct {
	Allowed_operations       []string
	Current_operations       map[string]string
	Currently_attached       bool
	Device                   string
	MAC                      string
	MAC_autogenerated        bool
	Metrics                  string
	MTU                      int
	Network                  string
	Other_config             map[string]string
	Qos_algorithm_params     map[string]string
	Qos_algorithm_type       string
	Qos_supported_algorithms []string
	Runtime_properties       map[string]string
	Status_code              int
	Status_detail            string
	Uuid                     string
	VM                       string
}

type PIF struct {
	Bond_master_of          []string
	Bond_slave_of           string
	Currently_attached      bool
	Device                  string
	Disallow_unplug         bool
	DNS                     string
	Gateway                 string
	Host                    string
	IP                      string
	Ip_configuration_mode   string
	MAC                     string
	Management              bool
	Metrics                 string
	MTU                     int
	Netmask                 string
	Network                 string
	Other_config            map[string]string
	Physical                bool
	Tunnel_access_PIF_of    string
	Tunnel_transport_PIF_of string
	Uuid                    string
	VLAN                    int
	VLAN_master_of          string
	VLAN_slave_of           []string
}

type Host struct {
	Address                           string
	Allowed_operations                []string
	Current_operations                map[string]string
	API_version_major                 int
	API_version_minor                 int
	API_version_vendor                string
	API_version_vendor_implementation map[string]string
	Bios_strings                      map[string]string
	Blobs                             map[string][]byte
	Capabilities                      []string
	Chipset_info                      map[string]string
	Cpu_configuration                 map[string]string
	Cpu_info                          map[string]string
	Crash_dump_sr                     string
	Crashdumps                        []string
	Edition                           string
	Enabled                           bool
	External_auth_configuration       map[string]string
	External_auth_type                string
	External_auth_service_name        string
	Ha_network_peers                  []string
	Ha_statefiles                     []string
	Host_CPUs                         []string
	Hostname                          string
	License_params                    map[string]string
	License_server                    map[string]string
	Local_cache_sr                    string
	Logging                           map[string]string
	Memory_overhead                   int
	Metrics                           string
	Name_description                  string
	Name_label                        string
	Other_config                      map[string]string
	Patches                           []string
	PBDs                              []string
	PCIs                              []string
	PGPUs                             []string
	PIFs                              []string
	Power_on_config                   map[string]string
	Resident_VMs                      []string
	Sched_policy                      string
	Software_version                  map[string]string
	Supported_bootloaders             []string
	Suspend_image_sr                  string
	Uuid                              string
	Tags                              []string
}
