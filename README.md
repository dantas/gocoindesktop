# gocoindesktop

This is my first Go project and it is still in progress.

The idea is to have a desktop app that displays the values of cryptocurrencies, can be minimized to systray and notify the client when the coin is about to enter/leave a specified range of values.

I'm using three different libraries to reach this goal, two of which I had to modify to make them work together. These modified versions are available here in my github.
- Fyne: for GUI and notifications. I modified it to execute an external function in its event loop.
- systray: Systray library. Built using GTK, I modified it to expose a single iteration of its event loop. 
- colly: for scrapping HTML

I thought about using Wire for dependency injection but since the application is supposed to be simple perhaps its better to create the composition root manually. More on that later.

Develop is on master because I'm the only sheriff on this town
