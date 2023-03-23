## Github
- Authentication

**SSH**
| Steps: create a new ssh key (public/private) and config the SSH config file to indicate the identity file to use
1. Open terminal and go to root folder
```bash
cd ~
```

2. Create and open .ssh folder (this contains the SSH keys and a config file)
```bash
mkdir .ssh && cd .ssh
```

3. Create a new SSH key
```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

4. Configure the SSH config file
```bash
touch config
```

5. Open the config file and add the following
```bash
Host *
  AddKeysToAgent yes
  UseKeychain yes
  IdentityFile ~/.ssh/id_rsa
```

<br>

**Github**
| Steps: add the SSH key (public) to your Github account
1. Copy the SSH key
```bash
pbcopy < ~/.ssh/id_rsa.pub
```

2. Go to your Github account and add the SSH key

<br>

**Test**
| Steps: We test the SSH connection to Github (the identity file is specified in the config file)
1. Test the SSH connection
```bash
ssh -T git@github.com
```

In order to test other identities, you can use the following command
```bash
ssh -i ~/.ssh/key -T git@github.com
```

<br>

Up to now you have configured the SSH key (public / private), indicate ssh which identity file to use and added the SSH key (public) to your Github account.
We can use an optional service to manage the SSH keys (cache) and avoid to enter the phrase each time we use the SSH key.

**SSH-Agent**
| Steps: initialize ssh-agent as a service of ssh keys management (cache) (optional)
1. Start the ssh-agent in the background
```bash
eval "$(ssh-agent -s)"
```

2. Add your SSH private key to the ssh-agent
```bash
ssh-add -K ~/.ssh/id_rsa
```

<br>

---

## Git
in progress...