import type { Post } from "./post";
import type { User } from "./user_types";

export class Favourite {
    public id: string;
    public post: Post;
    public user: User;
}