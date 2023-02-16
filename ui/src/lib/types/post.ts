import type { User } from "./user_types";

export class Post {
    public id: string;
    public createdAt: Date;
    public updatedAt: Date;
    public caption: string;
    public images: string[];
    public poster: User;
    public location: string;

    constructor(id: string, createdAt: Date, updatedAt: Date, caption: string, images: string[], poster: User, location: string) {
        this.id = id;
        this.createdAt = createdAt;
        this.updatedAt = updatedAt;
        this.caption = caption;
        this.images = images;
        this.poster = poster;
        this.location = location;
    }
}