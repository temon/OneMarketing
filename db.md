### Subscription Plan Table
- id int unique id of the subscription plan
- name varchar(255) name of the subscription plan
- description varchar(255) description of the subscription plan
- price decimal(10,2) price of the subscription plan
- billing_cycle int billing cycle in months
- trial_period int trial period in days
- status enum('active', 'inactive') status of the subscription plan
- created_at datetime timestamp of when the subscription plan was created
- updated_at datetime timestamp of when the subscription plan was last updated

### Plan Feature Table
- id int unique id of the plan feature
- name varchar(255) name of the plan feature
- description varchar(255) description of the plan feature
- code varchar(255) unique code of the plan feature
- type enum('limit', 'unlimited') type of the plan feature
- created_at datetime timestamp of when the plan feature was created
- updated_at datetime timestamp of when the plan feature was last updated

### Subscription Plan Feature Table
- id int unique id of the subscription plan feature
- subscription_plan_id int foreign key of the subscription plan
- plan_feature_id int foreign key of the plan feature
- limit_type enum('day', 'week', 'month', 'year') type of the plan feature limit
- limit_amount int amount of the plan feature limit
- created_at datetime timestamp of when the subscription plan feature was created
- updated_at datetime timestamp of when the subscription plan feature was last updated

### User Table
- id int unique id of the user
- name varchar(255) name of the user
- email varchar(255) email address of the user
- password varchar(255) hashed password of the user
- created_at datetime timestamp of when the user was created
- updated_at datetime timestamp of when the user was last updated

### User Subscription Table
- id int unique id of the user subscription
- user_id int foreign key of the user
- subscription_plan_id int foreign key of the subscription plan
- payment_id int foreign key of the payment
- start_date date date of when the subscription started
- end_date date date of when the subscription will end
- status enum('active', 'inactive') status of the user subscription
- created_at datetime timestamp of when the user subscription was created
- updated_at datetime timestamp of when the user subscription was last updated

### Payment Table
Field Name Data Type Description
- id int unique id of the payment
- user_id int foreign key of the user
- amount decimal(10,2) amount of the payment
- currency varchar(255) currency of the payment
- payment_date datetime date and time of when the payment was made
- status enum('paid', 'unpaid') status of the payment
- created_at datetime timestamp of when the payment was created
- updated_at datetime timestamp of when the payment was last updated

### API Usage Table
Field Name Data Type Description
- id int unique id of the API usage record
- user_subscription_id int foreign key of the user subscription
- api_endpoint varchar(255) endpoint of the API that was used
- request_count int number of requests made to the API endpoint
- usage_date date date of when the API was used
- created_at datetime timestamp of when the API usage record was created
- updated_at datetime timestamp of when the API usage record was last updated

### Campaign Table:
- campaign_id (int): primary key for the campaign
- user_id (int): foreign key to the user table
- data_source_id (int): foreign key to the data source table
- name (string): name of the campaign
- start_date (date): start date of the campaign
- end_date (date): end date of the campaign
- status (string): status of the campaign (e.g. active, paused, completed)
- created_at (timestamp): timestamp of when the campaign was created
- updated_at (timestamp): timestamp of when the campaign was last updated

### Campaign Performance Table:
- campaign_performance_id (int): primary key for the campaign performance
- campaign_id (int): foreign key to the campaign table
- metric_id (int): foreign key to the metric table
- value (float): value of the metric for the campaign
- date (date): date the metric was collected for the campaign
- created_at (timestamp): timestamp of when the campaign performance was created
- updated_at (timestamp): timestamp of when the campaign performance was last updated

### Channel Table:
- channel_id (int): primary key for the channel
- user_id (int): foreign key to the user table
- data_source_id (int): foreign key to the data source table
- name (string): name of the channel (e.g. Facebook, Twitter, Google Search)
- created_at (timestamp): timestamp of when the channel was created
- updated_at (timestamp): timestamp of when the channel was last updated

### Channel Performance Table:
- channel_performance_id (int): primary key for the channel performance
- channel_id (int): foreign key to the channel table
- metric_id (int): foreign key to the metric table
- value (float): value of the metric for the channel
- date (date): date the metric was collected for the channel
- created_at (timestamp): timestamp of when the channel performance was created
- updated_at (timestamp): timestamp of when the channel performance was last updated

### Ad Table:
- ad_id (int): primary key for the ad
- user_id (int): foreign key to the user table
- data_source_id (int): foreign key to the data source table
- campaign_id (int): foreign key to the campaign table
- name (string): name of the ad
- ad_type (string): type of the ad (e.g. display, search, video)
- ad_format (string): format of the ad (e.g. banner, text, pre-roll)
- created_at (timestamp): timestamp of when the ad was created
- updated_at (timestamp): timestamp of when the ad was last updated

### Ad Performance Table:
- ad_performance_id (int): primary key for the ad performance
- ad_id (int): foreign key to the ad table
- metric_id (int): foreign key to the metric table
- value (float): value of the metric for the ad
- date (date): date the metric was collected for the ad
- created_at (timestamp): timestamp of when the ad performance was created
- updated_at (timestamp): timestamp of when the ad performance was last updated