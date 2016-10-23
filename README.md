#Scraper development preperation by Leon
Dependencies needed:

1. For Golang html package:
go get golang.org/x/net/html

2. For mongodb go drive, two package needed

go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson

3. For scheduler need cron installed

go get github.com/robfig/cron

4. Google Stock api and description
 http://www.quantatrisk.com/2015/05/07/hacking-google-finance-in-pre-market-trading-python/

1)http://www.google.com/finance/info?client=ig&q=2388

2)http://www.google.com/finance/info?infotype=infoquoteall&q=2388,0911

3)http://www.google.com/finance/getprices?q=2388&i=60&p=1d&f=d,c,h,l,o,v


t 	Ticker
e 	Exchange
l 	Last Price
ltt 	Last Trade Time
l 	Price
lt 	Last Trade Time Formatted
lt_dts 	Last Trade Date/Time
c 	Change
cp 	Change Percentage
el 	After Hours Last Price
elt 	After Hours Last Trade Time Formatted
div 	Dividend
yld 	Dividend Yield

"price" - Realtime price quote, delayed by up to 20 minutes.
"priceopen" - The price as of market open.
"high" - The current day's high price.
"low" - The current day's low price.
"volume" - The current day's trading volume.
"marketcap" - The market capitalization of the stock.
"tradetime" - The time of the last trade.
"datadelay" - How far delayed the realtime data is.
"volumeavg" - The average daily trading volume.
"pe" - The price/earnings ratio.
"eps" - The earnings per share.
"high52" - The 52-week high price.
"low52" - The 52-week low price.
"change" - The price change since the previous trading day's close.
"beta" - The beta value.
"changepct" - The percentage change in price since the previous trading day's close.
"closeyest" - The previous day's closing price.
"shares" - The number of outstanding shares.
"currency" - The currency in which the security is priced.


5 google spreadsheet auto refresh

6 for ta-lib
if cannt not run
export LD_LIBRARY_PATH=/usr/lib:$LD_LIBRARY_PATH to fix  libta_lib.so.0 missing problem


