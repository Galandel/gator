GATOR - RSS feed aggregator

Requirements:
 - Go 
 - Postgres

To Install:
go install https://github.com/Galandel/gator.git

You need to create a .gatorconfig.json file in your home with the following information:
{"db_url":"postgres://"dbusername":"dbpassword"@"DBhost":"dbport"/gator?sslmode=disable","current_user_name":""}


Commands:
- gator register "username"       # creates user
- gator login "username"          # log user in
- gator addfeed "title" "url"     # add RSS feed to pull
- gator follow "url"              # follow a feed 
- gator following                 # Shows feeds user is following
- gator unfollow "url"            # Unfollow a feed
- gator agg "time: 1m"            # Pulls in Posts from feed
- gator browse "optiional: time"  # Shows posts in publish order newer first
