package main

import (
	"fmt"
	toogoodtogo "github.com/simonmartyr/toogoodtogogo"
	"log"
	"time"
)

type PickupInfo struct {
	StoreName  string
	ItemName   string
	PickupTime string
}

func ParseItems(items []*toogoodtogo.Item) *[]PickupInfo {
	var processed []PickupInfo
	for _, item := range items {
		pickupItem := PickupInfo{
			StoreName: item.Store.StoreName,
			ItemName:  item.Item.Name,
		}
		start, startErr := time.Parse(time.RFC3339, item.PickupInterval.Start)
		end, endErr := time.Parse(time.RFC3339, item.PickupInterval.End)
		timezone, timeZoneErr := time.LoadLocation(item.Store.StoreTimeZone)
		if startErr != nil || endErr != nil || timeZoneErr != nil {
			pickupItem.PickupTime = "unknown"
			log.Println(fmt.Sprintf("Date time error for: %s", item.Item.Name))
		} else {
			pickupItem.PickupTime = fmt.Sprintf("from: %s until %s",
				start.In(timezone).Format("15:04"),
				end.In(timezone).Format("15:04"),
			)
		}
		processed = append(processed, pickupItem)
	}
	return &processed
}

const emailTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Too Good To Go Availability</title>
    <style>
        /* Reset some default styles for better consistency */
        body, p, table {
            margin: 0;
            padding: 0;
        }

        /* Add a background color and padding to the body */
        body {
            background-color: #f0f0f0;
            padding: 20px;
        }

        /* Style the email container */
        .email-container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            padding: 20px;
            border-radius: 10px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        }

        /* Style the header */
        h1 {
            color: #333;
            font-size: 24px;
            margin-bottom: 20px;
        }

        /* Style the table */
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 20px;
        }

        table, th, td {
            border: 1px solid #ddd;
        }

        th, td {
            padding: 12px;
            text-align: left;
        }

        /* Style table header */
        th {
            background-color: #333;
            color: #fff;
        }

        /* Style table rows alternately */
        tr:nth-child(even) {
            background-color: #f2f2f2;
        }

        /* Style links */
        a {
            color: #007bff;
            text-decoration: none;
        }

        a:hover {
            text-decoration: underline;
        }

        /* Add spacing between elements */
        .spacer {
            margin-bottom: 20px;
        }

    </style>
</head>
<body>
<div class="email-container">
    <h1>Too Good To Go Availability</h1>

    <p>Dear User</p>

    <p>The following Too Good To Go Items are available:</p>

    <table>
        <tr>
            <th>Store Name</th>
            <th>Item Name</th>
            <th>Pickup Time</th>
        </tr>
        {{range .Items}}
        <tr>
            <td>{{.StoreName}}</td>
            <td>{{.ItemName}}</td>
            <td>{{.PickupTime}}</td>
        </tr>
        {{end}}
    </table>

    <div class="spacer"></div>

    <p>Thank you.</p>

    <p>Kind regards,<br>Raspberry Pi</p>
</div>
</body>
</html>
`
