# gocoindesktop

This is my first Go project and it is still in progress.

The idea is to have a desktop app that displays the values of cryptocurrencies, can be minimized to systray and notify the client when the coin is about to enter/leave a specified range of values.

I'm using three different libraries to reach this goal.
- Fyne: for GUI and notifications.
- systray: Systray library. 
- colly: for scrapping HTML

I thought about using Wire for dependency injection but since the application is supposed to be simple perhaps its better to create the composition root manually. More on that later.

Development is on master because I'm the only sheriff in this town
