import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;
import java.util.ArrayList;
import java.util.List;
import java.util.Comparator;
import java.util.Collections;

class Tournament {
    private Map<String, Team> teams = new HashMap<String, Team>();
    
    String printTable() {
        var table = "Team                           | MP |  W |  D |  L |  P\n";
        var teams = new ArrayList<Team>(this.teams.values());
        Collections.sort(teams, (team1, team2) -> {
            if (team1.points() == team2.points()) {
            return team1.name.compareTo(team2.name);
            }
            return team1.points() < team2.points() ? 1 : -1;
        });
        for (var team : teams) {
            table += String.format("%-30s | %2s | %2s | %2s | %2s | %2s\n",
                                  team.name,
                                  team.totalGames(),
                                  team.wins,
                                  team.draws,
                                  team.losses,
                                  team.points());
        }
        
        return table;
    }

    void applyResults(String resultString) {
        for (var line : resultString.split("\n")) {
            var match = line.split(";");
            var team1 = teams.computeIfAbsent(match[0], k -> new Team(k));
            var team2 = teams.computeIfAbsent(match[1], k -> new Team(k));
            addResults(team1, team2, match[2]);
        }
    }

    // mirrors the result input line
    void addResults(Team team1, Team team2, String result) {
        switch (result) {
            case "win": 
                team1.wins++;
                team2.losses++;
                break;
            case "draw": 
                team1.draws++;
                team2.draws++;
                break;
            case "loss": 
                team1.losses++;
                team2.wins++;
                break;
        }
    }
}

class Team {
    final String name;
    int wins;
    int losses;
    int draws;

    Team(String name) {
        this.name = name;
    }

    int points() {
        return wins * 3 + draws;
    }

    int totalGames() {
        return wins + losses + draws;
    }
}
