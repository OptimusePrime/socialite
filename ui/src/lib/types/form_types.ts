import type { IconDefinition } from "@fortawesome/free-brands-svg-icons";

export class InputLabelPairProps {
    placeholder?: string;
    label: string;
    type?: string;
    size?: "sm" | "md" | "lg" = "md";
    value?: string;
    color?: "base" | "green" | "red" = "base";
    inputClass?: string;
    labelClass?: string;
    helperText?: string;
    icon: string | IconDefinition;
}
