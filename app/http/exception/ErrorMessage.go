package exception

type ErrorMessage struct {
  Error bool      `json:"error"`
  Code int        `json:"code"`          
  Message string  `json:"message"`
  Details string  `json:"details"`
  Origin string   `json:"origin"`
}


