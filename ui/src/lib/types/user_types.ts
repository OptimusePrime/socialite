/*
ID        uuid.UUID `json:"id,omitempty"`
CreatedAt time.Time `json:"created_at,omitempty"`
UpdatedAt time.Time `json:"updated_at,omitempty"`
Username  string    `json:"username,omitempty"`
Email     string    `json:"email,omitempty"`
Name      string    `json:"name,omitempty"`
Password  string    `json:"password,omitempty"`
BirthDate time.Time `json:"birthDate,omitempty"`
Avatar    string    `json:"avatar,omitempty"`
Biography string    `json:"biography,omitempty"`
Gender    string    `json:"gender,omitempty"`
*/

export class User {
    public id: string;
    public username: string;
    public email: string;
    public name: string;
    public password: string;
    public birthDate: Date;
    public avatar: string;
    public biography: string;
    public gender: string;
    public pronouns: string;

}





