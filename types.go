package xapi

import (
	"time"
)

type Session struct {
	AuthUserName     string            `xmlrpc:"Auth_user_name"`
	AuthUserSid      string            `xmlrpc:"Auth_user_sid"`
	IsLocalSuperuser bool              `xmlrpc:"Is_local_superuser"`
	LastActive       time.Time         `xmlrpc:"Last_active"`
	OtherConfig      map[string]string `xmlrpc:"Other_config"`
	Parent           string
	Pool             bool
	RbacPermissions  []string `xmlrpc:"Rbac_permissions"`
	Subject          string
	Tasks            []string
	ThisHost         string    `xmlrpc:"This_host"`
	ThisUser         string    `xmlrpc:"This_user"`
	UUID             string    `xmlrpc:"Uuid"`
	ValidationTime   time.Time `xmlrpc:"Validation_time"`
}

type VDI struct {
	AllowCaching        bool              `xmlrpc:"Allow_caching"`
	AllowedOperations   []string          `xmlrpc:"Allowed_operations"`
	CrashDumps          []string          `xmlrpc:"Crash_dumps"`
	CurrentOperations   map[string]string `xmlrpc:"Current_operations"`
	IsAsnapshot         bool              `xmlrpc:"Is_a_snapshot"`
	Location            string
	Managed             bool
	MetadataLatest      bool `xmlrpc:"Metadata_latest"`
	Missing             bool
	NameDescription     string            `xmlrpc:"Name_description"`
	NameLabel           string            `xmlrpc:"Name_label"`
	OnBoot              string            `xmlrpc:"On_boot"`
	OtherConfig         map[string]string `xmlrpc:"Other_config"`
	Parent              string
	PhyscialUtilisation int  `xmlrpc:"Physcial_utilisation"`
	ReadOnly            bool `xmlrpc:"Read_only"`
	Sharable            bool
	SmConfig            map[string]string `xmlrpc:"Sm_config"`
	SnapshotOf          []string          `xmlrpc:"Snapshot_of"`
	SR                  []string
	StorageLock         bool `xmlrpc:"Storage_lock"`
	Tags                string
	Type                string `json:"type"`
	UUID                string `json:"uuid"`
	VDBs                []string
	VirtualSize         int               `xmlrpc:"Virtual_size"`
	XenstoreData        map[string]string `xmlrpc:"Xenstore_data"`
}

type VDB struct {
	AllowedOperations      []string          `xmlrpc:"Allowed_operations"`
	CurrentOperations      map[string]string `xmlrpc:"Current_operations"`
	Bootable               bool
	CurrentlyAttached      bool `xmlrpc:"Currently_attached"`
	Device                 string
	Empty                  bool
	Metrics                string
	Mode                   string
	OtherConfig            map[string]string `xmlrpc:"Other_config"`
	QosAlgorithmParams     map[string]string `xmlrpc:"Qos_algorithm_params"`
	QosAlgorithmType       string            `xmlrpc:"Qos_algorithm_type"`
	QosSupportedAlgorithms []string          `xmlrpc:"Qos_supported_algorithms"`
	RuntimeProperties      map[string]string `xmlrpc:"Runtime_properties"`
	StatusCode             int               `xmlrpc:"Status_code"`
	StatusDetail           string            `xmlrpc:"Status_detail"`
	UUID                   string            `json:"uuid"`
	VDI                    string
	VM                     string
	StorageLock            bool   `xmlrpc:"Storage_lock"`
	Type                   string `json:"type"`
	Unpluggable            bool
	Userdevice             string
}

type VM struct {
	ActionsAfterCrash       string `xmlrpc:"Actions_after_crash"`
	ActionsAfterReboot      string `xmlrpc:"Actions_after_reboot"`
	ActionsAfterShutdown    string `xmlrpc:"Actions_after_shutdown"`
	Affinity                string
	AllowedOperations       []string `xmlrpc:"Allowed_operations"`
	Appliance               string
	AttachedPCIs            []string          `xmlrpc:"Attached_PCIs"`
	BiosStrings             map[string]string `xmlrpc:"Bios_strings"`
	Blobs                   map[string][]byte
	BlockedOperations       map[string]string `xmlrpc:"Blocked_operations"`
	Children                []string
	Consoles                []string
	CrashDumps              []string          `xmlrpc:"Crash_dumps"`
	CurrentOperations       map[string]string `xmlrpc:"Current_operations"`
	Domarch                 string
	Domid                   int
	GuestMetrics            string            `xmlrpc:"Guest_metrics"`
	HaAlwaysRun             bool              `xmlrpc:"Ha_always_run"`
	HaRestartPriority       string            `xmlrpc:"Ha_restart_priority"`
	HVMBootParams           map[string]string `xmlrpc:"HVM_boot_params"`
	HVMBootPolicy           string            `xmlrpc:"HVM_boot_policy"`
	HVMShadowMultiplier     float64           `xmlrpc:"HVM_shadow_multiplier"`
	IsASnapshot             bool              `xmlrpc:"Is_a_snapshot"`
	IsATemplate             bool              `xmlrpc:"Is_a_template"`
	IsControlDomain         bool              `xmlrpc:"Is_control_domain"`
	IsSnapshotFromVmpp      bool              `xmlrpc:"Is_snapshot_from_vmpp"`
	LastBootCPUFlags        map[string]string `xmlrpc:"Last_boot_CPU_flags"`
	LastBootedRecord        string            `xmlrpc:"Last_booted_record"`
	MemoryDynamicMax        int               `xmlrpc:"Memory_dynamic_max"`
	MemoryDynamicMin        int               `xmlrpc:"Memory_dynamic_min"`
	MemoryOverhead          int               `xmlrpc:"Memory_overhead"`
	MemoryStaticMax         int               `xmlrpc:"Memory_static_max"`
	MemoryStaticMin         int               `xmlrpc:"Memory_static_min"`
	MemoryTarget            int               `xmlrpc:"Memory_target"`
	Metrics                 string
	NameDescription         string `xmlrpc:"Name_description"`
	NameLabel               string `xmlrpc:"Name_label"`
	Order                   int
	OtherConfig             map[string]string `xmlrpc:"Other_config"`
	Parent                  string
	PCIBus                  string `xmlrpc:"PCI_bus"`
	Platform                map[string]string
	PowerStatus             string `xmlrpc:"Power_status"`
	ProtectionPolicy        string `xmlrpc:"Protection_policy"`
	PVArgs                  string `xmlrpc:"PV_args"`
	PVBootloader            string `xmlrpc:"PV_bootloader"`
	PVBootloaderArgs        string `xmlrpc:"PV_bootloader_args"`
	PVKernel                string `xmlrpc:"PV_kernel"`
	PVLegacyArgs            string `xmlrpc:"PV_legacy_args"`
	PVRamdisk               string `xmlrpc:"PV_ramdisk"`
	Recommendations         string
	ResidentOn              string            `xmlrpc:"Resident_on"`
	ShutdownDelay           int               `xmlrpc:"Shutdown_delay"`
	SnapshotInfo            map[string]string `xmlrpc:"Snapshot_info"`
	SnapshotMetadata        string            `xmlrpc:"Snapshot_metadata"`
	SnapshotTime            time.Time         `xmlrpc:"Snapshot_time"`
	StartDelay              int               `xmlrpc:"Start_delay"`
	SuspendSR               string            `xmlrpc:"Suspend_SR"`
	SuspendVDI              string            `xmlrpc:"Suspend_VDI"`
	Tags                    []string
	TransportableSnapshotID string `xmlrpc:"Transportable_snapshot_id"`
	UserVersion             int    `xmlrpc:"User_version"`
	UUID                    string `json:"uuid"`
	VBDs                    []string
	VCPUsAtStartup          int               `xmlrpc:"VCPUs_at_startup"`
	VCPUsMax                int               `xmlrpc:"VCPUs_max"`
	VCPUsParams             map[string]string `xmlrpc:"VCPUs_params"`
	Version                 int
	VGPUs                   []string
	VIFs                    []string
	VTPM                    []string
	XenstoreData            map[string]string `xmlrpc:"Xenstore_data"`
}

type Event struct {
	Class     string
	ID        int    `xmlrpc:"id"`
	ObjUUID   string `xmlrpc:"Obj_uuid"`
	Operation string
	Ref       string
	Datetime  time.Time
}

type VIF struct {
	AllowedOperations      []string          `xmlrpc:"Allowed_operations"`
	CurrentOperations      map[string]string `xmlrpc:"Current_operations"`
	CurrentlyAttached      bool              `xmlrpc:"Currently_attached"`
	Device                 string
	MAC                    string
	MACAutogenerated       bool `xmlrpc:"MAC_autogenerated"`
	Metrics                string
	MTU                    int
	Network                string
	OtherConfig            map[string]string `xmlrpc:"Other_config"`
	QosAlgorithmParams     map[string]string `xmlrpc:"Qos_algorithm_params"`
	QosAlgorithmType       string            `xmlrpc:"Qos_algorithm_type"`
	QosSupportedAlgorithms []string          `xmlrpc:"Qos_supported_algorithms"`
	RuntimeProperties      map[string]string `xmlrpc:"Runtime_properties"`
	StatusCode             int               `xmlrpc:"Status_code"`
	StatusDetail           string            `xmlrpc:"Status_detail"`
	UUID                   string            `json:"uuid"`
	VM                     string
}

type PIF struct {
	BondMasterOf         []string `json:"Bond_master_of"`
	BondSlaveOf          string   `json:"Bond_slave_of"`
	CurrentlyAttached    bool     `json:"Currently_attached"`
	Device               string
	DisallowUnplug       bool `json:"Disallow_unplug"`
	DNS                  string
	Gateway              string
	Host                 string
	IP                   string
	IPConfigurationMode  string `json:"Ip_configuration_mode"`
	MAC                  string
	Management           bool
	Metrics              string
	MTU                  int
	Netmask              string
	Network              string
	OtherConfig          map[string]string `json:"Other_config"`
	Physical             bool
	TunnelAccessPIFOf    string `json:"Tunnel_access_PIF_of"`
	TunnelTransportPIFOf string `json:"Tunnel_transport_PIF_of"`
	UUID                 string `json:"uuid"`
	VLAN                 int
	VLANMasterOf         string   `json:"VLAN_master_of"`
	VLANSlaveOf          []string `json:"VLAN_slave_of"`
}

type Host struct {
	Address                        string
	AllowedOperations              []string          `json:"Allowed_operations"`
	CurrentOperations              map[string]string `json:"Current_operations"`
	APIVersionMajor                int               `json:"API_version_major"`
	APIVersionMinor                int               `json:"API_version_minor"`
	APIVersionVendor               string            `json:"API_version_vendor"`
	APIVersionVendorImplementation map[string]string `json:"API_version_vendor_implementation"`
	BiosStrings                    map[string]string `json:"Bios_strings"`
	Blobs                          map[string][]byte
	Capabilities                   []string
	ChipsetInfo                    map[string]string `json:"Chipset_info"`
	CPUConfiguration               map[string]string `json:"Cpu_configuration"`
	CPUInfo                        map[string]string `json:"Cpu_info"`
	CrashDumpSR                    string            `json:"Crash_dump_sr"`
	Crashdumps                     []string
	Edition                        string
	Enabled                        bool
	ExternalAuthConfiguration      map[string]string `json:"External_auth_configuration"`
	ExternalAuthType               string            `json:"External_auth_type"`
	ExternalAuthServiceName        string            `json:"External_auth_service_name"`
	HaNetworkPeers                 []string          `json:"Ha_network_peers"`
	HaStatefiles                   []string          `json:"Ha_statefiles"`
	HostCPUs                       []string          `json:"Host_CPUs"`
	Hostname                       string
	LicenseParams                  map[string]string `json:"License_params"`
	LicenseServer                  map[string]string `json:"License_server"`
	LocalCacheSR                   string            `json:"Local_cache_sr"`
	Logging                        map[string]string
	MemoryOverhead                 int `json:"Memory_overhead"`
	Metrics                        string
	NameDescription                string            `json:"Name_description"`
	NameLabel                      string            `json:"Name_label"`
	OtherConfig                    map[string]string `json:"Other_config"`
	Patches                        []string
	PBDs                           []string
	PCIs                           []string
	PGPUs                          []string
	PIFs                           []string
	PowerOnConfig                  map[string]string `json:"Power_on_config"`
	ResidentVMs                    []string          `json:"Resident_VMs"`
	SchedPolicy                    string            `json:"Sched_policy"`
	SoftwareVersion                map[string]string `json:"Software_version"`
	SupportedBootloaders           []string          `json:"Supported_bootloaders"`
	SuspendImageSR                 string            `json:"Suspend_image_sr"`
	UUID                           string            `json:"uuid"`
	Tags                           []string
}
