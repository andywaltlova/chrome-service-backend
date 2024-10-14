package models_test

import (
	"reflect"
	"testing"

	"github.com/RedHatInsights/chrome-service-backend/rest/models"
	"github.com/RedHatInsights/chrome-service-backend/rest/service"
	"github.com/stretchr/testify/assert"
)

const (
	validEncodedTemplate   = "eyJjcmVhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsInVwZGF0ZWRBdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZGVsZXRlZEF0IjpudWxsLCJ1c2VySWRlbnRpdHlJRCI6MCwiZGVmYXVsdCI6ZmFsc2UsIlRlbXBsYXRlQmFzZSI6eyJuYW1lIjoibGFuZGluZ1BhZ2UiLCJkaXNwbGF5TmFtZSI6IkxhbmRpbmcgUGFnZSJ9LCJ0ZW1wbGF0ZUNvbmZpZyI6eyJzbSI6W3siaCI6MSwiaSI6IkxhcmdlV2lkZ2V0I2x3MSIsInciOjEsIngiOjAsInkiOjAsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6IkxhcmdlV2lkZ2V0I2x3MiIsInciOjEsIngiOjAsInkiOjEsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6IkxhcmdlV2lkZ2V0I2x3MyIsInciOjEsIngiOjAsInkiOjIsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6Ik1lZGl1bVdpZGdldCNtdzEiLCJ3IjoxLCJ4IjoxLCJ5IjoyLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJTbWFsbFdpZGdldCNzdzEiLCJ3IjoxLCJ4IjoxLCJ5IjowLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJTbWFsbFdpZGdldCNzdzIiLCJ3IjoxLCJ4IjoxLCJ5IjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfV0sIm1kIjpbeyJoIjoxLCJpIjoiTGFyZ2VXaWRnZXQjbHcxIiwidyI6MSwieCI6MCwieSI6MCwibWF4SCI6NCwibWluSCI6MSwidGl0bGUiOiJXaWRnZXQgMSIsInN0YXRpYyI6dHJ1ZX0seyJoIjoxLCJpIjoiTGFyZ2VXaWRnZXQjbHcyIiwidyI6MSwieCI6MCwieSI6MSwibWF4SCI6NCwibWluSCI6MSwidGl0bGUiOiJXaWRnZXQgMSIsInN0YXRpYyI6dHJ1ZX0seyJoIjoxLCJpIjoiTGFyZ2VXaWRnZXQjbHczIiwidyI6MSwieCI6MCwieSI6MiwibWF4SCI6NCwibWluSCI6MSwidGl0bGUiOiJXaWRnZXQgMSIsInN0YXRpYyI6dHJ1ZX0seyJoIjoxLCJpIjoiTWVkaXVtV2lkZ2V0I213MSIsInciOjEsIngiOjIsInkiOjIsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6IlNtYWxsV2lkZ2V0I3N3MSIsInciOjEsIngiOjIsInkiOjAsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6IlNtYWxsV2lkZ2V0I3N3MiIsInciOjEsIngiOjIsInkiOjEsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9XSwibGciOlt7ImgiOjEsImkiOiJMYXJnZVdpZGdldCNsdzEiLCJ3IjoxLCJ4IjowLCJ5IjowLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJMYXJnZVdpZGdldCNsdzIiLCJ3IjoxLCJ4IjowLCJ5IjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJMYXJnZVdpZGdldCNsdzMiLCJ3IjoxLCJ4IjowLCJ5IjoyLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJNZWRpdW1XaWRnZXQjbXcxIiwidyI6MSwieCI6MywieSI6MiwibWF4SCI6NCwibWluSCI6MSwidGl0bGUiOiJXaWRnZXQgMSIsInN0YXRpYyI6dHJ1ZX0seyJoIjoxLCJpIjoiU21hbGxXaWRnZXQjc3cxIiwidyI6MSwieCI6MywieSI6MCwibWF4SCI6NCwibWluSCI6MSwidGl0bGUiOiJXaWRnZXQgMSIsInN0YXRpYyI6dHJ1ZX0seyJoIjoxLCJpIjoiU21hbGxXaWRnZXQjc3cyIiwidyI6MSwieCI6MywieSI6MSwibWF4SCI6NCwibWluSCI6MSwidGl0bGUiOiJXaWRnZXQgMSIsInN0YXRpYyI6dHJ1ZX1dLCJ4bCI6W3siaCI6MSwiaSI6IkxhcmdlV2lkZ2V0I2x3MSIsInciOjEsIngiOjAsInkiOjAsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6IkxhcmdlV2lkZ2V0I2x3MiIsInciOjEsIngiOjAsInkiOjEsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6IkxhcmdlV2lkZ2V0I2x3MyIsInciOjEsIngiOjAsInkiOjIsIm1heEgiOjQsIm1pbkgiOjEsInRpdGxlIjoiV2lkZ2V0IDEiLCJzdGF0aWMiOnRydWV9LHsiaCI6MSwiaSI6Ik1lZGl1bVdpZGdldCNtdzEiLCJ3IjoxLCJ4Ijo0LCJ5IjoyLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJTbWFsbFdpZGdldCNzdzEiLCJ3IjoxLCJ4Ijo0LCJ5IjowLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfSx7ImgiOjEsImkiOiJTbWFsbFdpZGdldCNzdzIiLCJ3IjoxLCJ4Ijo0LCJ5IjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJ0aXRsZSI6IldpZGdldCAxIiwic3RhdGljIjp0cnVlfV19fQo="
	invalidEncodedTemplate = "eyJjcmVhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsInVwZGF0ZWRBdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZGVsZXRlZEF0IjpudWxsLCJ1c2VySWRlbnRpdHlJRCI6MCwiZGVmYXVsdCI6ZmFsc2UsIlRlbXBsYXRlQmFzZSI6eyJuYW1lIjoiIiwiZGlzcGxheU5hbWUiOiJ0ZXN0In0sInRlbXBsYXRlQ29uZmlnIjp7InNtIjpbeyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IkxhcmdlV2lkZ2V0I2x3MSIsIngiOjAsInkiOjAsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IkxhcmdlV2lkZ2V0I2x3MiIsIngiOjAsInkiOjEsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IkxhcmdlV2lkZ2V0I2x3MyIsIngiOjAsInkiOjIsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6Ik1lZGl1bVdpZGdldCNtdzEiLCJ4IjoxLCJ5IjoyLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJTbWFsbFdpZGdldCNzdzEiLCJ4IjoxLCJ5IjowLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJTbWFsbFdpZGdldCNzdzIiLCJ4IjoxLCJ5IjoxLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9XSwibWQiOlt7InRpdGxlIjoiV2lkZ2V0IDEiLCJpIjoiTGFyZ2VXaWRnZXQjbHcxIiwieCI6MCwieSI6MCwidyI6MSwiaCI6MSwibWF4SCI6NCwibWluSCI6MSwic3RhdGljIjp0cnVlfSx7InRpdGxlIjoiV2lkZ2V0IDEiLCJpIjoiTGFyZ2VXaWRnZXQjbHcyIiwieCI6MCwieSI6MSwidyI6MSwiaCI6MSwibWF4SCI6NCwibWluSCI6MSwic3RhdGljIjp0cnVlfSx7InRpdGxlIjoiV2lkZ2V0IDEiLCJpIjoiTGFyZ2VXaWRnZXQjbHczIiwieCI6MCwieSI6MiwidyI6MSwiaCI6MSwibWF4SCI6NCwibWluSCI6MSwic3RhdGljIjp0cnVlfSx7InRpdGxlIjoiV2lkZ2V0IDEiLCJpIjoiTWVkaXVtV2lkZ2V0I213MSIsIngiOjIsInkiOjIsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IlNtYWxsV2lkZ2V0I3N3MSIsIngiOjIsInkiOjAsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IlNtYWxsV2lkZ2V0I3N3MiIsIngiOjIsInkiOjEsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX1dLCJsZyI6W3sidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJMYXJnZVdpZGdldCNsdzEiLCJ4IjowLCJ5IjowLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJMYXJnZVdpZGdldCNsdzIiLCJ4IjowLCJ5IjoxLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJMYXJnZVdpZGdldCNsdzMiLCJ4IjowLCJ5IjoyLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJNZWRpdW1XaWRnZXQjbXcxIiwieCI6MywieSI6MiwidyI6MSwiaCI6MSwibWF4SCI6NCwibWluSCI6MSwic3RhdGljIjp0cnVlfSx7InRpdGxlIjoiV2lkZ2V0IDEiLCJpIjoiU21hbGxXaWRnZXQjc3cxIiwieCI6MywieSI6MCwidyI6MSwiaCI6MSwibWF4SCI6NCwibWluSCI6MSwic3RhdGljIjp0cnVlfSx7InRpdGxlIjoiV2lkZ2V0IDEiLCJpIjoiU21hbGxXaWRnZXQjc3cyIiwieCI6MywieSI6MSwidyI6MSwiaCI6MSwibWF4SCI6NCwibWluSCI6MSwic3RhdGljIjp0cnVlfV0sInhsIjpbeyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IkxhcmdlV2lkZ2V0I2x3MSIsIngiOjAsInkiOjAsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IkxhcmdlV2lkZ2V0I2x3MiIsIngiOjAsInkiOjEsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6IkxhcmdlV2lkZ2V0I2x3MyIsIngiOjAsInkiOjIsInciOjEsImgiOjEsIm1heEgiOjQsIm1pbkgiOjEsInN0YXRpYyI6dHJ1ZX0seyJ0aXRsZSI6IldpZGdldCAxIiwiaSI6Ik1lZGl1bVdpZGdldCNtdzEiLCJ4Ijo0LCJ5IjoyLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJTbWFsbFdpZGdldCNzdzEiLCJ4Ijo0LCJ5IjowLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9LHsidGl0bGUiOiJXaWRnZXQgMSIsImkiOiJTbWFsbFdpZGdldCNzdzIiLCJ4Ijo0LCJ5IjoxLCJ3IjoxLCJoIjoxLCJtYXhIIjo0LCJtaW5IIjoxLCJzdGF0aWMiOnRydWV9XX19Cg=="
	invalidEncodedStruct   = "ewogICAgImZvbyI6ICJiYXIiCn0="
	invalidEncodedJson     = "ewogICAgImZvbyI6ICJiYXIKfQ=="
	matchTestString        = "eyJjcmVhdGVkQXQiOiIwMDAxLTAxLTAxVDAwOjAwOjAwWiIsInVwZGF0ZWRBdCI6IjAwMDEtMDEtMDFUMDA6MDA6MDBaIiwiZGVsZXRlZEF0IjpudWxsLCJ1c2VySWRlbnRpdHlJRCI6MCwiZGVmYXVsdCI6ZmFsc2UsIlRlbXBsYXRlQmFzZSI6eyJuYW1lIjoidGVzdCIsImRpc3BsYXlOYW1lIjoidGVzdCJ9LCJ0ZW1wbGF0ZUNvbmZpZyI6eyJzbSI6bnVsbCwibWQiOm51bGwsImxnIjpudWxsLCJ4bCI6bnVsbH19Cg=="
)

var invalidBaseNameTemplate = models.DashboardTemplate{
	TemplateBase: models.DashboardTemplateBase{
		Name:        "",
		DisplayName: "test",
	},
	TemplateConfig: service.BaseTemplates["landingPage"].TemplateConfig,
}

var invalidBaseDisplayNameTemplate = models.DashboardTemplate{
	Default: false,
	TemplateBase: models.DashboardTemplateBase{
		Name:        "test",
		DisplayName: "",
	},
	TemplateConfig: service.BaseTemplates["landingPage"].TemplateConfig,
}

var validTemplate = models.DashboardTemplate{
	Default: false,
	TemplateBase: models.DashboardTemplateBase{
		Name:        "test",
		DisplayName: "test",
	},
	TemplateConfig: service.BaseTemplates["landingPage"].TemplateConfig,
}

func TestDashboardTemplateValidation(t *testing.T) {
	itct := service.BaseTemplates["landingPage"].TemplateConfig
	invalidData := []models.GridItem{
		{
			BaseWidgetDimensions: models.BaseWidgetDimensions{
				Width:     1,
				Height:    1,
				MaxHeight: 1,
				MinHeight: 1,
			},
			Title: "Widget 1",
			ID:    "LargeWidget_lw1",
			X:     2,
			Y:     0,
		},
	}
	invalidData[0].X = 2
	itct.Sm = service.ConvertToJson(invalidData)
	invalidTemplateConfigTemplate := models.DashboardTemplate{
		Default: false,
		TemplateBase: models.DashboardTemplateBase{
			Name:        "test",
			DisplayName: "test",
		},
		TemplateConfig: itct,
	}

	type testCase struct {
		Name         string
		Input        models.DashboardTemplate
		ErrorMessage string
		IsValid      bool
	}

	cases := []testCase{
		{
			Name:         "invalidBaseNameTemplate",
			Input:        invalidBaseNameTemplate,
			ErrorMessage: "invalid template name",
		},
		{
			Name:         "invalidBaseDisplayNameTemplate",
			Input:        invalidBaseDisplayNameTemplate,
			ErrorMessage: "invalid template display name",
		},
		{
			Name:         "invalidTemplateConfigTemplate",
			Input:        invalidTemplateConfigTemplate,
			ErrorMessage: "invalid grid item, layout variant sm, coordinate X must be less than 1, current value is 2",
		}, {
			Name:         "validTemplate",
			Input:        validTemplate,
			ErrorMessage: "",
			IsValid:      true,
		},
	}

	for _, c := range cases {
		err := c.Input.IsValid()
		if c.IsValid {
			assert.Nil(t, err)
			continue
		}
		assert.NotNil(t, err)
		assert.Equal(t, c.ErrorMessage, err.Error())
	}
}

func TestDashboardTemplateEncoding(t *testing.T) {
	validTemplate := models.DashboardTemplate{
		Default: false,
		TemplateBase: models.DashboardTemplateBase{
			Name:        "test",
			DisplayName: "test",
		},
		TemplateConfig: service.BaseTemplates["landingPage"].TemplateConfig,
	}

	t.Run("Should encode template", func(t *testing.T) {
		encoded, err := validTemplate.EncodeBase64()
		assert.Nil(t, err)
		assert.Equal(t, encoded, matchTestString)
	})

	t.Run("Should decode template", func(t *testing.T) {
		decoded, err := models.DecodeDashboardBase64(matchTestString)
		assert.Nil(t, err)
		assert.Equal(t, reflect.TypeOf(decoded), reflect.TypeOf(models.DashboardTemplate{}))
	})

	t.Run("Should fail to decode invalid template", func(t *testing.T) {
		_, err := models.DecodeDashboardBase64(invalidEncodedTemplate)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid template name", err.Error())
	})

	t.Run("Should fail to decode invalid structure", func(t *testing.T) {
		_, err := models.DecodeDashboardBase64(invalidEncodedStruct)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid template name", err.Error())
	})

	t.Run("Should fail to decode invalid json", func(t *testing.T) {
		_, err := models.DecodeDashboardBase64(invalidEncodedJson)
		assert.NotNil(t, err)
		assert.Equal(t, "invalid character '\\n' in string literal", err.Error())
	})
}
