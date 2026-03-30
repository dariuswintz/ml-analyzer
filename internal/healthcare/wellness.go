package healthcare

type SimpleWellnessAdvisor struct {}

func (s *SimpleWellnessAdvisor) RecommendExercise() string {
	return "I recommend a daily walk or jog for at least 30 minutes." 
}

func (s *SimpleWellnessAdvisor) RecommendDiet() string {
	return "A balanced diet rich in fruits, vegetables, and whole grains is ideal." 
}

func (s *SimpleWellnessAdvisor) CreateWellnessPlan() string {
	return "A typical wellness plan includes exercise, diet management, and regular check-ups." 
}