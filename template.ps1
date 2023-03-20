# GO BASES
New-Item -Path .\ -Name "01-go-bases" -ItemType Directory
for ($i = 1; $i -le 5; $i++) {
    # directories
    New-Item -Path .\01-go-bases\ -Name "clase-0$i-TM" -ItemType Directory
    New-Item -Path .\01-go-bases\clase-0$i-TM\ -Name "live" -ItemType Directory
    New-Item -Path .\01-go-bases\clase-0$i-TM\ -Name "practice" -ItemType Directory
    # files
    New-Item -Path .\01-go-bases\clase-0$i-TM\live\ -Name ".gitkeep" -ItemType File
    New-Item -Path .\01-go-bases\clase-0$i-TM\practice\ -Name ".gitkeep" -ItemType File
}

# GO WEB
New-Item -Path .\ -Name "02-go-web" -ItemType Directory
# for loop to add class folders
for ($i = 1; $i -le 5; $i++) {
    # directories
    New-Item -Path .\02-go-web\ -Name "clase-0$i-TM" -ItemType Directory
    New-Item -Path .\02-go-web\clase-0$i-TM\ -Name "live" -ItemType Directory
    New-Item -Path .\02-go-web\clase-0$i-TM\ -Name "practice" -ItemType Directory
    # files
    New-Item -Path .\02-go-web\clase-0$i-TM\live\ -Name ".gitkeep" -ItemType File
    New-Item -Path .\02-go-web\clase-0$i-TM\practice\ -Name ".gitkeep" -ItemType File
}

# GO DATABASE
New-Item -Path .\ -Name "03-go-database" -ItemType Directory
# for loop to add class folders
for ($i = 1; $i -le 3; $i++) {
    # directories
    New-Item -Path .\03-go-database\ -Name "clase-0$i-TM" -ItemType Directory
    New-Item -Path .\03-go-database\clase-0$i-TM\ -Name "live" -ItemType Directory
    New-Item -Path .\03-go-database\clase-0$i-TM\ -Name "practice" -ItemType Directory
    # files
    New-Item -Path .\03-go-database\clase-0$i-TM\live\ -Name ".gitkeep" -ItemType File
    New-Item -Path .\03-go-database\clase-0$i-TM\practice\ -Name ".gitkeep" -ItemType File
}

# GO STORAGE
New-Item -Path .\ -Name "04-go-storage" -ItemType Directory
# for loop to add class folders
for ($i = 1; $i -le 2; $i++) {
    # directories
    New-Item -Path .\04-go-storage\ -Name "clase-0$i-TM" -ItemType Directory
    New-Item -Path .\04-go-storage\clase-0$i-TM\ -Name "live" -ItemType Directory
    New-Item -Path .\04-go-storage\clase-0$i-TM\ -Name "practice" -ItemType Directory
    # files
    New-Item -Path .\04-go-storage\clase-0$i-TM\live\ -Name ".gitkeep" -ItemType File
    New-Item -Path .\04-go-storage\clase-0$i-TM\practice\ -Name ".gitkeep" -ItemType File
}

# GO TESTING
New-Item -Path .\ -Name "05-go-testing" -ItemType Directory
# for loop to add class folders
for ($i = 1; $i -le 4; $i++) {
    # directories
    New-Item -Path .\05-go-testing\ -Name "clase-0$i-TM" -ItemType Directory
    New-Item -Path .\05-go-testing\clase-0$i-TM\ -Name "live" -ItemType Directory
    New-Item -Path .\05-go-testing\clase-0$i-TM\ -Name "practice" -ItemType Directory
    # files
    New-Item -Path .\05-go-testing\clase-0$i-TM\live\ -Name ".gitkeep" -ItemType File
    New-Item -Path .\05-go-testing\clase-0$i-TM\practice\ -Name ".gitkeep" -ItemType File
}