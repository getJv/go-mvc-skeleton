package exception

type ErrorMessage struct {
  Error bool      `json:"error" example:"true"`
  Code int        `json:"code" example:"401"`          
  Message string  `json:"message" example:"Something happened..."`
  Details string  `json:"details" example:"Error description"`
  Origin string   `json:"origin" example:"File and Line of given error"`
}


