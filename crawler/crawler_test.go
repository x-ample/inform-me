package crawler

import "testing"

func TestCrawler(teste *testing.T)  {
    cases := []struct {
        expected []string
    }

    for case_key, case_value : = range cases {
        actual_result := Crawler()
        if actual_result != case_value.expected {
            teste.Errorf("%s, is not %s ", actual_result, case_value.expected)
        }
    }

}
