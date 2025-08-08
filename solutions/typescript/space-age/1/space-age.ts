export default class SpaceAge {
  readonly seconds: number;

  constructor(seconds: number) {
    this.seconds = seconds;
  }

  private toEarthYears(orbitalPeriod: number) {
    return Math.round((this.seconds / (orbitalPeriod * 31557600)) * 100) / 100;
  }

  onEarth() {
    return this.toEarthYears(1);
  }
  onMercury() {
    return this.toEarthYears(0.2408467);
  }
  onVenus() {
    return this.toEarthYears(0.61519726);
  }
  onMars() {
    return this.toEarthYears(1.8808158);
  }
  onJupiter() {
    return this.toEarthYears(11.862615);
  }
  onSaturn() {
    return this.toEarthYears(29.447498);
  }
  onUranus() {
    return this.toEarthYears(84.016846);
  }
  onNeptune() {
    return this.toEarthYears(164.79132);
  }
}
