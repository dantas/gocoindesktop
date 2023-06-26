# Go Coin Desktop

This is my first Go project. It is developed and tested in Linux, with no guarantees if it works on other OSes.

It is a a desktop app that displays the values of cryptocurrencies, can be minimized to systray and notify the client when the coin is about to enter/leave a specified range of value.

## Screenshots

![Coin Tab](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/coins.jpg)

![Settings Tab](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/settings.jpg)

![Alert Notification](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/alert.jpg)

![System Tray](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/systemtray.jpg)

## Development Notes

I'm using three different libraries to achieve this goal.
- Fyne: for GUI and notifications.
- systray: Systray library. 
- colly: for scrapping HTML

No library for dependency injection. I'm manually constructing the composition root.

Development is on master because I'm the only sheriff in this town

### Build Tags

- **pt**: Replace english with portuguese (pt-br).
- **testui**: Test UI without any domain and IO.

### Build system requirements

Ensure the following packages are installed to successfully build the app on ubuntu:  
    ```sudo apt-get install pkg-config xorg-dev libayatana-appindicator3-dev```

