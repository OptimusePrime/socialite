export function validateEmail(email: string): boolean {
    return (
        email
            .toLowerCase()
            .match(
                /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
            ) !== null && email.length < 64
    );
}

export function validateUsername(username: string): boolean {
    return username.length > 3 && username.length < 16;
}

export function validateGender(gender: string): boolean {
    return gender.length < 16;
}

export function validateName(name: string): boolean {
    return name.length > 3 && name.length < 16;
}

export function validatePassword(password: string): boolean {
    return (
        new RegExp(/^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,32}$/).test(password)
    );
}

export function generatePassword(length: number): string {
    const charset =
		"!\"#&'()*,-./:;?@[]ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let password = "";
    for (let i = 0; i < length; i++) {
        password += charset[Math.floor(Math.random() * charset.length)];
    }

    return password;
}
