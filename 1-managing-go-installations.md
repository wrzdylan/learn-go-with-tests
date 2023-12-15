# Gestion de Go

## Installer Go dans Debian

- Télécharger Go : `curl -O -L https://go.dev/dl/go1.21.5.linux-amd64.tar.gz`         

> -O: Indique à curl de sauvegarder le fichier téléchargé avec son nom d'origine dans le répertoire actuel.          
> -L: Permet à curl de suivre les redirections si le lien conduit vers une redirection vers un autre emplacement.         

- Extraire les fichiers : `sudo tar -C /usr/local/ -xzf go1.21.5.linux-amd64.tar.gz`     

- Include go dans l'environnement en ajoutant la ligne suivante dans le fichier ~/.profile : `export PATH=$PATH:/usr/local/go/bin`     


## Installer plusieurs versions de Go avec gvm

- https://github.com/moovweb/gvm         

- `sudo apt-get install bison`      
- `bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)`                     

(En cas d'erreur : GVM couldn't find hexdump) - `sudo apt-get install bsdmainutils`       

- `gvm install go1.11.4 --binary`
- `gvm use go1.11.4`
