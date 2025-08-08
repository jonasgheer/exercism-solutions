export default class SpaceAge {
  constructor(readonly seconds: number) {}

  private toEarthYears = (orbitalPeriod: number) => () =>
    Math.round((this.seconds / (orbitalPeriod * 31557600)) * 100) / 100;

  onEarth = this.toEarthYears(1);
  onMercury = this.toEarthYears(0.2408467);
  onVenus = this.toEarthYears(0.61519726);
  onMars = this.toEarthYears(1.8808158);
  onJupiter = this.toEarthYears(11.862615);
  onSaturn = this.toEarthYears(29.447498);
  onUranus = this.toEarthYears(84.016846);
  onNeptune = this.toEarthYears(164.79132);
}
