# Set The Siliq Lora Package Forwarder

This package forwoarder fork from RAK <https://github.com/RAKWireless/rak_common_for_gateway>, but add install.go to finish some setting

## install

1. clone this Repository from github
2. execute install
3. execute start.sh to have the local_conf.json
4. enable lora_pkt_fwd service
5. set up the parameter

```bash
git clone https://github.com/livinghuang/lora_pkg_fwd.git
chmod 755 install
./install
./set_eui.sh
systemctl start lora_pkt_fwd.service
system enable lora_pkt_fwd.service
nano /opt/siliq/global_conf.json
```
