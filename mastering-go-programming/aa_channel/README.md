# Challenges

Based on video here: https://www.youtube.com/watch?v=f6kdp27TYZs

* Create a generator function which returns a channel, and launches a go routtine which eternally writes to the returned channel
* Do this for two channels
* Create a fanin function which makes it easier to consume from two channels at once
* Make the initial generator function use a control channel, so we can synchronise order in which they read
