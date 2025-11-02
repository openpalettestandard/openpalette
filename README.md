# OpenPalette Specification

**An Open-Source Design Language & Palette Framework**

This is work in progress, the tooling isn't finished and the specification and documentation isn't finished either. Subject to change, but the goal is to have a standardized format and toolchain to create color palette organizations to bring themes to various applications.

## 1. Overview

OpenPalette is a standardized, open-source design language and palette framework that provides a unified structure for creating consistent, accessible, and interoperable color palettes. It bridges the gap between existing design systems by establishing common standards for accent colors, semantic hierarchy, and terminal integration.

## 2. Core Principles

- **Semantic Structure**: Colors are defined by their semantic role, not just their appearance
- **Universal Compatibility**: 14-accent system ensures compatibility across design frameworks
- **Terminal Integration**: First-class support for terminal and ANSI color mapping
- **Accessibility**: Built-in contrast and accessibility considerations
- **Open Standards**: Community-driven, open-source specification

## 3. Palette Structure

### 3.1 Accent Colors (14 Required)

Every OpenPalette-compliant palette MUST define exactly 14 accent colors:

| Name       | Purpose               | Notes                             |
| ---------- | --------------------- | --------------------------------- |
| `red`      | Primary red accent    | Error states, destructive actions |
| `maroon`   | Deep red variant      | Dark red tones                    |
| `pink`     | Red-purple midpoint   | Accent variation                  |
| `orange`   | Warm accent           | Warning states, highlights        |
| `yellow`   | Bright accent         | Attention, notifications          |
| `green`    | Primary green accent  | Success states, positive actions  |
| `teal`     | Green-cyan midpoint   | Cool accent variation             |
| `cyan`     | Cool accent           | Information, links                |
| `sky`      | Light blue variant    | Secondary information             |
| `blue`     | Primary blue accent   | Primary actions, links            |
| `sapphire` | Blue variant          | Accent variation                  |
| `purple`   | Primary purple accent | Special states, creativity        |
| `mauve`    | Purple variant        | Accent variation                  |
| `lavender` | Light purple          | Subtle accents                    |

### 3.2 Semantic Elements (12 Required)

Every OpenPalette-compliant palette MUST define exactly 12 semantic elements in hierarchical order:

| Element    | Order | Purpose                               |
| ---------- | ----- | ------------------------------------- |
| `text`     | 0     | Primary text color                    |
| `subtext1` | 1     | Secondary text, labels                |
| `subtext0` | 2     | Tertiary text, captions, placeholders |
| `overlay2` | 3     | Highest overlay level, modals         |
| `overlay1` | 4     | Mid-level overlays, tooltips          |
| `overlay0` | 5     | Low-level overlays, hover states      |
| `surface2` | 6     | Elevated surfaces, cards              |
| `surface1` | 7     | Mid-level surfaces, containers        |
| `surface0` | 8     | Base surfaces, input fields           |
| `base`     | 9     | Default background                    |
| `mantle`   | 10    | Secondary background                  |
| `crust`    | 11    | Border color, deepest background      |

## 4. Data Format Specification

### 4.1 Required Fields

Each palette MUST include:

```json
{
  "name": "string",
  "version": "semver",
  "dark": boolean,
  "colors": {
    // 14 accent colors + 12 semantic elements
  }
}
```

### 4.2 Color Object Format

Each color MUST include:

```json
{
  "name": "string",
  "order": number,
  "hex": "#000000",
  "rgb": {
    "r": 0,
    "g": 0, 
    "b": 0
  },
  "hsl": {
    "h": 0,
    "s": 0.0,
    "l": 0.0
  },
  "accent": boolean
}
```

### 4.3 Complete Palette Example

```json
{
  "name": "Example Palette",
  "version": "1.0.0",
  "dark": true,
  "colors": {
    "red": {
      "name": "Red",
      "order": 0,
      "hex": "#e86671",
      "rgb": { "r": 232, "g": 102, "b": 113 },
      "hsl": { "h": 355, "s": 0.74, "l": 0.65 },
      "accent": true
    },
    // ... all 26 colors required
  }
}
```

## 5. ANSI Terminal Mapping
OpenPalette defines standard ANSI color mapping for terminal compatibility

## 6. Accessibility Requirements

### 6.1 Contrast Standards

- text on base: MUST meet WCAG AA (4.5:1)
- subtext1 on base: MUST meet WCAG AA (4.5:1)
- subtext0 on base: SHOULD meet WCAG AA (4.5:1)
- Accent colors on base: SHOULD meet WCAG AA when used for text

### 6.2 Color Differentiation

- All accent colors MUST be distinguishable from each other
- Semantic hierarchy MUST be visually apparent
- Palette MUST work for common color vision differences

## 7. Validation Rules

### 7.1 Required Elements

- MUST contain exactly 14 accent colors
- MUST contain exactly 12 semantic elements
- MUST include all required color object fields
- MUST use valid hex color codes
- MUST include valid RGB values (0-255)
- MUST include valid HSL values (h: 0-360, s/l: 0-1)

### 7.2 Naming Conventions

- Color names MUST use lowercase
- No spaces or special characters in color names
- Palette name SHOULD be descriptive and unique

## 8. Implementation Guidelines

### 8.1 File Structure
```text
palette-name/
├── palette.json          # Core palette definition
├── README.md            # Documentation
├── LICENSE              # Open source license
└── ports/               # Platform-specific implementations
    ├── terminal/
    ├── vscode/
    ├── web/
    └── ...
```

## 8.2 Versioning

- Use semantic versioning (MAJOR.MINOR.PATCH)
- Major version changes for breaking color modifications
- Minor version for new features or non-breaking additions
- Patch version for fixes and adjustments

## 9. Community Standards

### 9.1 Contribution Guidelines

- All palettes MUST be open source
- MUST include proper attribution
 - SHOULD include usage examples
- SHOULD provide multiple format exports

### 9.2 Quality Standards

- MUST pass OpenPalette validation
- SHOULD include accessibility testing results
- SHOULD provide both light and dark variants
- MUST include comprehensive documentation

## 10. Compliance
A palette is OpenPalette-compliant if it:

- Contains exactly 26 colors (14 accents + 12 semantic)
- Uses the specified data format
- Meets accessibility requirements
- Passes validation rules
- Includes required metadata
