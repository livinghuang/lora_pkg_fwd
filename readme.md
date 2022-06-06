# Set The Siliq Lora Package Forwarder

This package forwoarder fork from RAK <https://github.com/RAKWireless/rak_common_for_gateway>, but add install.go to finish some setting. It already tested by Raspberry Pi3+ A.

## install

1. clone this Repository from github
2. execute install
3. enable lora_pkt_fwd service
4. set up the parameter

```bash
# if you use new install RPI, please open SPI interface  
sudo raspi-config
sudo apt-get update
sudo apt-get upgrade -y
sudo apt-get install git
mkdir lora
cd lora
git clone https://github.com/livinghuang/lora_pkg_fwd.git
cd lora_pkg_fwd/
# you could rebuild install by below command , but if you use rpi3 or rpi4 , just use install_rpi3 may workable
# go build install.go
# sudo chmod 755 install
sudo ./install_rpi3
sudo systemctl start lora_pkt_fwd.service
sudo systemctl status lora_pkt_fwd.service
sudo systemctl enable lora_pkt_fwd.service
sudo nano /opt/siliq/global_conf.json
```
