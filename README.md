#HubSpot

HubSpot API in Go.


#Examples
```golang
import(
  "github.com/Fyb3roptik/hubspot"
)

api_key := os.Getenv("HUBSPOT_API_KEY")

a := NewContact(api_key, "abhi@acksin.com")
a.Add("firstname", "Abhi")
a.Add("lastname", "Yerra")
a.Add("company", "Acksin")
a.Add("lifecyclestage", "opportunity")
a.Add("acksinsoftware", "opsZero")
resp := a.Publish()
if resp.Vid != 901 {
        t.Errorf("Failed to update contact")
}

d := NewDeal(api_key)
d.Associations.AssociatedVids = []int{resp.Vid}
d.Add("dealname", "Tim's Newer Deal")
d.Add("dealstage", "closedwon")
d.Add("closedate", Timestamp())
d.Add("amount", "60000")
d.Add("dealtype", "newbusiness")
d.Publish()

// Single Send Email has 2 property types. Contact and Custom, so we need to specify using the first param
e := hubspot.NewEmail(api_key, email_id, to_email_address)

// Adding a contact property
e.Add("contact", "firstname", "Jack")

// Adding a custom property
e.Add("custom", "some_custom_key", "some_custom_value")

e.Publish()
```
# Credit
This is based on the original library written by acksin
