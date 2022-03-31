package cert

import "testing"

func TestValidCertDate(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert Date doit être valid. err =%v", err)
	}
	if c == nil {
		t.Errorf("cert doit être une reference valid. obtenu=nil ")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Le nom du cours n'est pas valid. attendu = 'Golang', obtenu=  %v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Error("Error doit retourner un cours vide")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "iuoazhefrhnfvkqehglzerfgjelqaghrlefgjhnledrfghlerghlekrghlerkgekejgfldfghlazihezouhgfbjsvqz"
	_, err := New(course, "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("L’erreur doit être retournée sur un nom de cours trop long (cours=%s", course)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2018-05-31")
	if err == nil {
		t.Error("Error doit retourner un cours vide")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "iuoazhefrhnfvkqehglzerfgjelqaghrlefgjhnledrfghlerghlekrghlerkgekejgfldfghlazihezouhgfbjsvqz"
	_, err := New("Golang", name, "2018-05-31")
	if err == nil {
		t.Errorf("L’erreur doit être retournée sur un nom trop long (nom=%s", name)
	}
}
