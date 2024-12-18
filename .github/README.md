# Go Coin Desktop

This is a desktop app that displays the values of cryptocurrencies, can be minimized to systray and notifies the client when the coin is entering/leaving specified values.

It is developed and tested in Linux, with no guarantees if it works on other OSes.

**Status**: Completed

## Screenshots

![Coin Tab](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/coins.jpg)

![Settings Tab](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/settings.jpg)

![Alert Notification](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/alert.jpg)

![System Tray](https://raw.githubusercontent.com/dantas/gocoindesktop/master/.github/readmedia/systemtray.jpg)

# Developer Notes

I'm using two libraries to achieve this goal.
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
**Presenter** isolates the UI from the rest of the application. We can build the app in test ui mode (using *testui* flag) and develop/test the ui without having to worry about the rest of the code.  
Domain is tested in [domain_test.go](https://github.com/dantas/gocoindesktop/blob/master/domain/domain_test.go). These tests are integration tests that follow the [Big Bang Method](https://www.linkedin.com/advice/0/how-can-big-bang-integration-testing-save-time), but with external API and filesystem stubed out to make the tests predictable and independent of each other.
