# gocoindesktop

This is my first Go project and it is still in progress. It is developed and tested in Linux, with no guarantees if it will work on other OSes.

The idea is to have a desktop app that displays the values of cryptocurrencies, can be minimized to systray and notify the client when the coin is about to enter/leave a specified range of value.

I'm using three different libraries to achieve this goal.
- Fyne: for GUI and notifications.
- systray: Systray library. 
- colly: for scrapping HTML

No library for dependency injection. I'm manually constructing the composition root

Development is on master because I'm the only sheriff in this town

### Go build options

There is also an experimental command line output. Use build tag **cmd** to enable it.

### Build system requirements

Ensure the following packages are installed to successfully build the app on ubuntu:  
    ```sudo apt-get install pkg-config xorg-dev libayatana-appindicator3-dev```
