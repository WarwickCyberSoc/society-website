# Generates the HTML for team team.md

# Final entry must be "BREAK"
team = [
    ("Josh H.", "President", "joshh.jpg"),
    ("Alex M.", "Vice President", "alexm.jpg"),
    "BREAK",
    ("Adam B.", "Sr. Social Executive", "adamb.jpg"),
    ("Angus G.", "Sr. Social Executive", "angusg.jpg"),
    ("Joe B.", "Social Executive", "joeb.jpg"),
    ("Chris H.", "Social Executive", "chrish.jpg"),
    "BREAK",
    ("Julia L.", "Sr. Welfare Officer", "julial.jpg"),
    ("James J.", "Welfare Officer", "jamesj.jpg"),
    ("Freya T.", "Sr. Treasurer", "freyat.jpg"),
    ("Oscar C.", "Treasurer", "oscarc.jpg"),
    "BREAK",
    ("Drew G.", "Sr. External Comms Officer", "drewg.jpg"),
    ("Kunal S.", "External Comms Officer", "kunals.jpg"),
    ("Cam W.", "Technical Officer", "camw.jpg"),
    "BREAK",
]

output = ""
team_row_container = '<div class="team-row">{}</div>'
team_member = '<div class="team-member"><img src="/images/team/{image}"><p><strong>{name}</strong><br>{role}</p></div>'

temp_output = ""

for member in team:
    if member == "BREAK":
        output += team_row_container.format(temp_output)
        temp_output = ""
        continue

    temp_output += team_member.format(name=member[0], role=member[1], image=member[2])

print(output)
