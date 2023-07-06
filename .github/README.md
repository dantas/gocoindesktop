# Go Coin Desktop

This is a desktop app that displays the values of cryptocurrencies, can be minimized to systray and notify the client when the coin is about to enter/leave a specified range of value.

It is developed and tested in Linux, with no guarantees if it works on other OSes.

**Status**: Completed

## Screenshots

![Coin Tab](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/coins.jpg)

![Settings Tab](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/settings.jpg)

![Alert Notification](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/alert.jpg)

![System Tray](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/systemtray.jpg)

# Developer Notes

I'm using three different libraries to achieve this goal.
- Fyne: for GUI and notifications.
- systray: Systray library. 

No library for dependency injection. I'm manually constructing the composition root.

Development is on master because I'm the only sheriff in this town

## Build Tags

- **pt**: Replace english with portuguese (pt-br).
- **testui**: Test UI without any domain and IO.

## Build system requirements

Ensure the following packages are installed to successfully build the app on ubuntu:  
    ```sudo apt-get install pkg-config xorg-dev libayatana-appindicator3-dev```

## Architecture

Since this is my first Go project, I'm using the minimum amount of third party libraries.  
I purposely avoided using many Fyne features like data binding and storage since I'm learning/practicing Go and I want to code more, not less.  
This app uses **Clean Architecture**. Domain is isolated from external API and filesystem. **Dependency injection** is used to provide the appropriate dependencies.  
**Presenter** is isolated from the rest of the application, so that one can build the app (using *testui* flag) and develop/test the UI without having to worry about the rest of the code.
