- Format one device as a fat32 device
  mkfs.fat -F 32 {devicename}

  mkfs.fat -F 32 /dev/sdb


- Format one device as an ext4 device
  mkfs.ext4 {devicename}

  mkfs.ext4 /dev/sda


- Mount one of the devices

  mount {devicename} {mountpoint}


Run the test program on each device:

  go-fs-test {mountpoint}/file1.txt

  (unplug device randomly here)

  (plug device back in and check that data in the files is valid)
