# **DRAFT**

## mount RAM filesystem

```Bash
sudo mount -t ramfs ramfs /mnt/storage-ram
```

## helpful commands

### list info about disk

```Bash
sudo fdisk -l [ /dev/... ]
```

### dm status

```Bash
sudo cryptsetup status { mapper-name }
sudo integritysetup status { mapper-name }
```

## script

```Bash
# mount
sudo mkdir -pv /mnt/storage-ram/bck4
sudo chown -R vitalik:vitalik /mnt/storage-ram/bck4

_bck_dev=/dev/sdd1 \
    && _bck4_ikey=/mnt/storage-ram/bck4/bck4-ikey \
    && _bck4_header=/mnt/storage-ram/bck4/bck4-header \
    && _bck4_key=/mnt/storage-ram/bck4/bck4-key \
    && _bck4_imapper=bck4-int \
    && _bck4_emapper=bck4-enc \
    && _bck4_mount=/mnt/bck7
echo "dev: ${_bck_dev}" &&
    echo "ikey: ${_bck4_ikey}" &&
    echo "imapper: ${_bck4_imapper}" &&
    echo "header: ${_bck4_header}" &&
    echo "key: ${_bck4_key}" &&
    echo "emapper: ${_bck4_emapper}" &&
    echo "mount: ${_bck4_mount}"

# generate random data
sudo dd if=/dev/urandom bs=32 count=1 of=$_bck4_ikey
sudo dd if=/dev/urandom bs=512 count=1 of=$_bck4_key
sudo dd if=/dev/zero bs=16777216 count=1 of=$_bck4_header

# format and open integration layer
sudo integritysetup --debug format $_bck_dev --tag-size=32 --integrity=hmac-sha256 --integrity-key-file=$_bck4_ikey --integrity-key-size=32 --sector-size=4096
sudo integritysetup --debug open $_bck_dev $_bck4_imapper --integrity=hmac-sha256 --integrity-key-file=$_bck4_ikey --integrity-key-size=32

# format encryption layer
sudo cryptsetup --debug luksFormat --cipher=capi:xts(twofish)-essiv:sha256 --key-size=256 --pbkdf=argon2id --iter-time=16000 --hash=sha512 --label=bck4-vault --key-slot=4 --use-urandom --key-file=$_bck4_key --header=$_bck4_header --sector-size=4096 /dev/mapper/$_bck4_imapper
sudo cryptsetup --debug luksOpen --key-slot=4 --key-file=$_bck4_key --header=$_bck4_header /dev/mapper/$_bck4_imapper $_bck4_emapper

# make filesystem
sudo mkfs.ext4 -F -v /dev/mapper/$_bck4_emapper

# mount
sudo mkdir -pv $_bck4_mount
sudo mount -t ext4 /dev/mapper/$_bck4_emapper $_bck4_mount/

# encrypt credentials
gpg2 --verbose -u 069feac3229cbcf1 -r me@vitalik-malkin.email --encrypt $_bck4_ikey
gpg2 --verbose -u 069feac3229cbcf1 -r me@vitalik-malkin.email --encrypt $_bck4_header
gpg2 --verbose -u 069feac3229cbcf1 -r me@vitalik-malkin.email --encrypt $_bck4_key

# copy encrypted credentials to separate place
sudo mkdir -pv /media/vitalik/bck4-key/secrets/cryptsetup
cp $_bck4_ikey.gpg /media/vitalik/bck4-key/secrets/cryptsetup/
cp $_bck4_header.gpg /media/vitalik/bck4-key/secrets/cryptsetup/
cp $_bck4_key.gpg /media/vitalik/bck4-key/secrets/cryptsetup/

```
