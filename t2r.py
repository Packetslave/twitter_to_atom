#!env python

import datetime
import feedgenerator
import twitter

OUT = '/srv/www/example.com/public_html/feeds'
USERS = '/home/example/users.txt'

with open(USERS, 'r') as f:
  screen_names = f.readlines()

a = twitter.Api(
  consumer_key='MYKEY',
  consumer_secret='MYSECRET',
  access_token_key='MYTOKENKEY',
  access_token_secret='MYTOKENSECRET')

for user in (x.strip() for x in screen_names):
  feed = feedgenerator.Rss201rev2Feed(
    title=u"Tweets for %s" % user,
    link=u"http://twitter.com/%s" % user,
    description=u"Tweets for %s" % user,
    language=u"en")

  statuses = a.GetUserTimeline(screen_name=user)

  for status in statuses:
    pubdate = datetime.datetime.strptime(
        status.created_at, '%a %b %d %H:%M:%S +0000 %Y')

    link='http://twitter.com/%s/status/%s' % (user, status.id)

    feed.add_item(
      title=status.text,
      description=status.text,
      unique_id=link,
      link=link,
      pubdate=pubdate)

  with open('%s/%s.rss' % (OUT, user), 'w') as f:
    feed.write(f, 'utf-8')
