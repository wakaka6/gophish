package models

import (
	"fmt"

	check "gopkg.in/check.v1"
)

type mockTemplateContext struct {
	URL         string
	FromAddress string
}

func (m mockTemplateContext) getFromAddress() string {
	return m.FromAddress
}

func (m mockTemplateContext) getBaseURL() string {
	return m.URL
}

func (s *ModelsSuite) TestNewTemplateContext(c *check.C) {
	r := Result{
		BaseRecipient: BaseRecipient{
			FirstName: "Foo",
			LastName:  "Bar",
			Email:     "foo@bar.com",
		},
		RId: "1234567",
	}
	ctx := mockTemplateContext{
		URL:         "http://example.com",
		FromAddress: "From Address <from@example.com>",
	}
	expected := PhishingTemplateContext{
		URL:           fmt.Sprintf("%s?rid=%s", ctx.URL, r.RId),
		BaseURL:       ctx.URL,
		BaseRecipient: r.BaseRecipient,
		TrackingURL:   fmt.Sprintf("%s/track?rid=%s", ctx.URL, r.RId),
		QrcodeURL:     fmt.Sprintf("%s/qrcode?rid=%s", ctx.URL, r.RId),
		From:          "From Address",
		RId:           r.RId,
	}
	qrcodeImg := "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAACQElEQVR42tyZvZHrMAyEl6OAoUtQKSqNLI2luASFCjjEm11APruDJzA4j8nPwYH4WYAwrY7NbGCBq61q76P7SRJAe0dfG78Vu178vF7vo/tJFmAz/dOTmICBVWUUnSQDZAFsdu4DRX9TApOXWyZ00fUEkgHh1cCLIXxht/7Z+nb7ZwORo9Z2vUYrs577aEXB+5vEHg7EYtgO0A44rNDDO37W0wHd8GfpqNCdrfD7SALInUdjdmII1xNHB0PYL/pIApSpSOXHYQufHHXY2uxMAyxUe7vvclXJiI9X75YEgN8tudGKXdhdFLpR7CFARGAxOwGbUrPaywXce+60XjcXaA24PMgBwFNQmdVMpfKNRukuo1Dz5QDkyApEyQOGYzdWz9B/aQC6c1MKagv13M0r5i32kgDT9+TVFhZQhxJenQNQOhpN+qdHI7ag5lpaIQmwVB0NqjjYLkkipS3ePfuxFIAmIu7V3o6xylaZhmnLsgDy6i5lB6alfTRXEUUungQosTdrqAiGrX7jP80DxBiPFcd7bGXnm34IoHmA5pOdIuHNowA+qfjxgIU80AyEf6NfdmGUBojkquu+jWLqubyIJgHM9Z2/HVjEJv7GIkmAWEtir4RolxaCDwdyADGftImXppQsJYravr7E3uOB72nz/QYUM5Bb+aUA/O2AYXuYrDHa7dWXZkH5gHIHLK3BeLaRC1jKUS4mpGn1+qVX2hyAe7XPJ10Ztdj6a2GeD0SOKspR/kobb+6/qfi/Bv4FAAD//+exVsBW9qRIAAAAAElFTkSuQmCC"
	expected.Tracker = "<img alt='' style='display: none' src='" + expected.TrackingURL + "'/>"
	expected.Qrcode = "<img alt='' style='border: 5px solid #fff;' src='data:image/png;base64," + qrcodeImg + "'/>"
	expected.BaseQrcode = "data:image/png;base64," + qrcodeImg

	got, err := NewPhishingTemplateContext(ctx, r.BaseRecipient, r.RId)
	c.Assert(err, check.Equals, nil)
	c.Assert(got, check.DeepEquals, expected)
}
