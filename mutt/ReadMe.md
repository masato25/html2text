Setup Mutt

```
mkdir -p ~/.config/mutt

cp mailcap ~/.config/mutt
cp muttrc ~/.config/mutt
cp init.sh ~/.config/mutt

ln -s ~/.config/mutt/muttrc ~/.muttrc

cd ~/.config/mutt ./init.sh
```

* modified muttrc to setup your mail settings
* change html2text compiled path for mailcap
