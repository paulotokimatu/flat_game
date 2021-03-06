#version 410

in vec2 TexCoords;
out vec4 color;

uniform sampler2D tex;
uniform vec4 labelColor;

void main()
{    
  vec4 sampled = vec4(1.0, 1.0, 1.0, texture(tex, TexCoords).r);
  color = labelColor * sampled;
}
