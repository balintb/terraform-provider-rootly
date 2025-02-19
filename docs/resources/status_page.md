---
page_title: "Resource rootly_status_page - terraform-provider-rootly"
subcategory:
description: |-
    
---

# Resource (rootly_status_page)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `title` (String) The title of the status page

### Optional

- `allow_search_engine_index` (Boolean) Allow search engines to include your public status page in search results
- `authentication_enabled` (Boolean) Enable authentication
- `authentication_password` (String) Authentication password
- `description` (String) The description of the status page
- `enabled` (Boolean)
- `failure_message` (String) Message showing when at least one component is not operational
- `footer_color` (String) The color of the footer. Eg. "#1F2F41"
- `functionality_ids` (List of String) Functionalities attached to the status page
- `ga_tracking_id` (String) Google Analytics tracking ID
- `header_color` (String) The color of the header. Eg. "#0061F2"
- `public` (Boolean) Make the status page accessible to the public
- `public_description` (String) The public description of the status page
- `public_title` (String) The public title of the status page
- `service_ids` (List of String) Services attached to the status page
- `show_uptime` (Boolean) Show uptime
- `show_uptime_last_days` (Number) Show uptime over x days. Value must be one of `30`, `60`, `90`, `180`, `360`.
- `success_message` (String) Message showing when all components are operational
- `time_zone` (String) Status Page Timezone. Value must be one of `International Date Line West`, `American Samoa`, `Midway Island`, `Hawaii`, `Alaska`, `Pacific Time (US & Canada)`, `Tijuana`, `Arizona`, `Mazatlan`, `Mountain Time (US & Canada)`, `Central America`, `Central Time (US & Canada)`, `Chihuahua`, `Guadalajara`, `Mexico City`, `Monterrey`, `Saskatchewan`, `Bogota`, `Eastern Time (US & Canada)`, `Indiana (East)`, `Lima`, `Quito`, `Atlantic Time (Canada)`, `Caracas`, `Georgetown`, `La Paz`, `Puerto Rico`, `Santiago`, `Newfoundland`, `Brasilia`, `Buenos Aires`, `Montevideo`, `Greenland`, `Mid-Atlantic`, `Azores`, `Cape Verde Is.`, `Edinburgh`, `Lisbon`, `London`, `Monrovia`, `UTC`, `Amsterdam`, `Belgrade`, `Berlin`, `Bern`, `Bratislava`, `Brussels`, `Budapest`, `Casablanca`, `Copenhagen`, `Dublin`, `Ljubljana`, `Madrid`, `Paris`, `Prague`, `Rome`, `Sarajevo`, `Skopje`, `Stockholm`, `Vienna`, `Warsaw`, `West Central Africa`, `Zagreb`, `Zurich`, `Athens`, `Bucharest`, `Cairo`, `Harare`, `Helsinki`, `Jerusalem`, `Kaliningrad`, `Kyiv`, `Pretoria`, `Riga`, `Sofia`, `Tallinn`, `Vilnius`, `Baghdad`, `Istanbul`, `Kuwait`, `Minsk`, `Moscow`, `Nairobi`, `Riyadh`, `St. Petersburg`, `Volgograd`, `Tehran`, `Abu Dhabi`, `Baku`, `Muscat`, `Samara`, `Tbilisi`, `Yerevan`, `Kabul`, `Ekaterinburg`, `Islamabad`, `Karachi`, `Tashkent`, `Chennai`, `Kolkata`, `Mumbai`, `New Delhi`, `Sri Jayawardenepura`, `Kathmandu`, `Almaty`, `Astana`, `Dhaka`, `Urumqi`, `Rangoon`, `Bangkok`, `Hanoi`, `Jakarta`, `Krasnoyarsk`, `Novosibirsk`, `Beijing`, `Chongqing`, `Hong Kong`, `Irkutsk`, `Kuala Lumpur`, `Perth`, `Singapore`, `Taipei`, `Ulaanbaatar`, `Osaka`, `Sapporo`, `Seoul`, `Tokyo`, `Yakutsk`, `Adelaide`, `Darwin`, `Brisbane`, `Canberra`, `Guam`, `Hobart`, `Melbourne`, `Port Moresby`, `Sydney`, `Vladivostok`, `Magadan`, `New Caledonia`, `Solomon Is.`, `Srednekolymsk`, `Auckland`, `Fiji`, `Kamchatka`, `Marshall Is.`, `Wellington`, `Chatham Is.`, `Nuku'alofa`, `Samoa`, `Tokelau Is.`.
- `website_privacy_url` (String) Website Privacy URL
- `website_support_url` (String) Website Support URL
- `website_url` (String) Website URL

### Read-Only

- `id` (String) The ID of this resource.
