Practice of back of the envelop calculations
---

Here are just my back of the envelop calculations practices.

### How many read/write/search requests in Twitter?

#### Number of requests

- 300M twitter users in total.
- 100M active users
- 500M tweets per day.
  - Estimate 5 tweets/day by 1 active user.
- 5B read requests/day.
  - Estimate 5 times as much as 500M tweets.
- 200M search queries/day.

#### Size of the data

NOTE: [Refereced this site](https://github.com/donnemartin/system-design-primer/tree/master/solutions/system_design/twitter#calculate-usage)

- Size of tweets
  - tweet_id  8 bytes
  - user_id   32 bytes
  - text      140 bytes
  - media     10 KB average
  - TOTAL     around 10KB

- MEMO: The maximum limit of pictures/GIF size are 5MB for picttures. 5MB(from mob)/15MB(from twitter.com) for GIF.
- MEMO: 1 month ≒ 31 days ≒ 750 hours ≒ 2.5M seconds
